package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theofalso/timetrack/internal/session"
	"github.com/theofalso/timetrack/internal/store"
)

var startCmd = &cobra.Command{
	Use:   "start [project]",
	Short: "Start tracking a project",
	Long:  `Starts a new time session for the specified project.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		sessions, err := store.Load()
		if err != nil {
			return fmt.Errorf("error loading sessions: %w", err)
		}
		if len(sessions) > 0 {
			lastSession := sessions[len(sessions)-1]
			if lastSession.IsActive() {
				return fmt.Errorf("there is already an active session for project '%s'. Stop it first", lastSession.Project)
			}
		}
		newSession := session.NewSession(projectName)
		sessions = append(sessions, *newSession)

		if err := store.Save(sessions); err != nil {
			return fmt.Errorf("error saving session: %w", err)
		}

		fmt.Printf("Starting session for project '%s' at %s\n", newSession.Project, newSession.StartTime.Format("15:04:05"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
