package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/timperman/gobilize/mobilize"
)

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

		events, err := mobilize.ListEvents(mobilize.ListEventsRequest{
			OrganizationID: orgID,
			TimeslotStart:  fmt.Sprintf("gte_%d", startTime.Unix()),
			TimeslotEnd:    fmt.Sprintf("lte_%d", endTime.Unix()),
			ZipCode:        zip,
		})
		if err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "URL"})
		table.SetAutoWrapText(false)
		for _, event := range events {
			table.Append([]string{fmt.Sprintf("%d", event.ID), event.Title, event.BrowserURL})
		}
		table.Render()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
