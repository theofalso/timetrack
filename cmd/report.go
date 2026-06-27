package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/theofalso/timetrack/internal/store"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Show tracked time report",
	Long:  `Displays a summary of total tracked time grouped by project.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		sessions, err := store.Load()
		if err != nil {
			return fmt.Errorf("error loading sessions: %w", err)
		}

		if len(sessions) == 0 {
			fmt.Println("📭 No recorded sessions found.")
			return nil
		}

		// map to hold total durations per project
		projectTotals := make(map[string]time.Duration)

		for _, s := range sessions {
			projectTotals[s.Project] += s.Duration()
		}

		fmt.Println("Time Report by Project:")
		fmt.Println("---------------------------")

		for project, totalTime := range projectTotals {
			cleanTime := totalTime.Round(time.Second)
			if cleanTime >= time.Minute {
				cleanTime = totalTime.Round(time.Minute)
			}

			fmt.Printf("%s: %s\n", project, cleanTime)
		}
		fmt.Println("---------------------------")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
