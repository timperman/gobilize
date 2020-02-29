package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/timperman/gobilize/mobilize"
)

var start, zip string
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

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Date", "Title", "URL"})
		table.SetAutoMergeCells(true)
		table.SetAutoWrapText(false)
		for date, events := range eventsByDate {
			for _, event := range events {
				title := fmt.Sprintf("%s - %s", event.Timeslots[0].StartDate.Time().Format("3:04pm"), event.Title)
				table.Append([]string{date, title, event.BrowserURL})
			}
		}
		table.Render()
		return nil
	},
}

func init() {
	listCmd.Flags().StringVarP(&start, "start", "s", "", fmt.Sprintf("Event start time in %s format (default now)", time.RFC3339))
	listCmd.Flags().IntVarP(&days, "days", "d", 7, "Number of days since start")
	listCmd.Flags().IntVarP(&orgID, "org-id", "o", 1767, "Organization ID")
	listCmd.Flags().StringVarP(&zip, "zip-code", "z", "", "Events near ZIP code")
	listCmd.Flags().IntVar(&maxDistance, "max-dist", 0, "Maximum distance")

	rootCmd.AddCommand(listCmd)
}
