package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theofalso/timetrack/web"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server dashboard",
	Long:  `Launches a local HTTP server to view your time tracking statistics in the browser.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := web.StartServer("8080")
		if err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
