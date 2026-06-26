package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/theofalso/timetrack/internal/store"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the active session",
	Long:  `Stops the currently running time session and calculates the duration.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		sessions, err := store.Load()
		if err != nil {
			return fmt.Errorf("error loading sessions: %w", err)
		}

		// check if there is an active session
		if len(sessions) == 0 {
			return fmt.Errorf("there are no sessions recorded")
		}

		lastIndex := len(sessions) - 1
		if !sessions[lastIndex].IsActive() {
			return fmt.Errorf("there is no active session to stop")
		}

		// stop the active session
		sessions[lastIndex].EndTime = time.Now()
		duration := sessions[lastIndex].Duration()

		// save the updated sessions back to the store
		if err := store.Save(sessions); err != nil {
			return fmt.Errorf("error saving session: %w", err)
		}

		fmt.Printf("Stopped session for '%s'\n", sessions[lastIndex].Project)
		fmt.Printf("Duration: %s\n", duration.Round(time.Second))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
