package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/timperman/gobilize/handle"
)

var port string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		http.HandleFunc("/events", handle.RenderEvents)

		fmt.Printf("Starting server on %s...", port)
		return http.ListenAndServe(port, nil)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", ":8080", "HTTP port")

	rootCmd.AddCommand(serveCmd)
}
