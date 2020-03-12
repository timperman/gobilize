package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/timperman/gobilize/mobilize"
)

var start, zip, region, format string
var days, orgID, maxDistance int

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List events",
	RunE: func(cmd *cobra.Command, args []string) error {
		startTime := time.Now()
		var err error
		if start != "" {
			startTime, err = time.Parse(time.RFC3339, start)
			if err != nil {
				return err
			}
		}
		endTime := startTime.Add(time.Duration(days) * 24 * time.Hour)

		eventsByDate, err := mobilize.ListEventsByDate(mobilize.ListEventsRequest{
			OrganizationID: orgID,
			TimeslotStart:  fmt.Sprintf("gte_%d", startTime.Unix()),
			TimeslotEnd:    fmt.Sprintf("lte_%d", endTime.Unix()),
			ZipCode:        zip,
			MaxDistance:    maxDistance,
		})
		if err != nil {
			return err
		}

		switch format {
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Date", "Title", "URL"})
			table.SetAutoMergeCells(true)
			table.SetAutoWrapText(false)
			for date, events := range eventsByDate {
				for _, event := range events {
					if region == "" || event.Location.Region == region {
						title := fmt.Sprintf("%s - %s", event.Timeslots[0].StartDate.Time().Format("3:04pm"), event.Title)
						table.Append([]string{date, title, event.BrowserURL})
					}
				}
			}
			break
		case "csv":
			fmt.Println("date,time,title,venue,address,city,state,zip,timeslots,capacity,url")
			for date, events := range eventsByDate {
				for _, event := range events {
					if region == "" || event.Location.Region == region {
						addr := ""
						for _, line := range event.Location.AddressLines {
							if addr == "" {
								addr = fmt.Sprintf("%s", line)
							} else if line != "" {
								addr = fmt.Sprintf("%s, %s", addr, line)
							}
						}
						fmt.Printf("\"%s\",%s,\"%s\",\"%s\",\"%s\",%s,%s,%s,%d,%d,%s\n", date, event.Timeslots[0].StartDate.Time().Format("3:04pm"), event.Title, event.Location.Venue, addr, event.Location.Locality, event.Location.Region, event.Location.PostalCode, len(event.Timeslots), event.Timeslots[0].MaxAttendees, event.BrowserURL)
					}
				}
			}
		}
		return nil
	},
}

func init() {
	listCmd.Flags().StringVarP(&start, "start", "s", "", fmt.Sprintf("Event start time in %s format (default now)", time.RFC3339))
	listCmd.Flags().IntVarP(&days, "days", "d", 7, "Number of days since start")
	listCmd.Flags().IntVarP(&orgID, "org-id", "o", 1767, "Organization ID")
	listCmd.Flags().StringVarP(&zip, "zip-code", "z", "", "Events near ZIP code")
	listCmd.Flags().IntVar(&maxDistance, "max-dist", 0, "Maximum distance")
	listCmd.Flags().StringVarP(&region, "region", "r", "", "Region (state code) filter")
	listCmd.Flags().StringVarP(&format, "format", "f", "table", "Output format (table, csv, or html)")

	rootCmd.AddCommand(listCmd)
}
