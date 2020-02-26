package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var start, zip string
var days, orgID int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gobilize",
	Short: "A tool for interacting with MobilizeAmerica events",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&start, "start", "s", "", fmt.Sprintf("Event start time in %s format (default now)", time.RFC3339))
	rootCmd.PersistentFlags().IntVarP(&days, "days", "d", 7, "Number of days since start")
	rootCmd.PersistentFlags().IntVarP(&orgID, "org-id", "o", 1767, "Organization ID")
	rootCmd.PersistentFlags().StringVarP(&zip, "zip-code", "z", "", "Events near ZIP code")
}
