package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/theofalso/timetrack/internal/store"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show active session status",
	Long:  `Displays the currently active session and its elapsed time in real-time.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// load sessions from the store
		sessions, err := store.Load()
		if err != nil {
			return fmt.Errorf("error loading sessions: %w", err)
		}

		// this verify if there is an active session, if not, it will print a message and exit
		if len(sessions) == 0 {
			fmt.Println("💤 No active sessions.")
			return nil
		}
		lastSession := sessions[len(sessions)-1]
		if !lastSession.IsActive() {
			fmt.Println("💤 No active sessions.")
			return nil
		}

		fmt.Printf("tracking project: %s\n", lastSession.Project)
		fmt.Println("Press Ctrl+C to exit.")

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		// 5. infinite loop to update the elapsed time every second
		for {
			select {
			case <-ticker.C:
				duration := time.Since(lastSession.StartTime).Round(time.Second)
				fmt.Printf("\r⏱️ Elapsed time: %s", duration)

			case <-sigChan: // exit with control c
				fmt.Println("\nExiting status...")
				return nil
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
