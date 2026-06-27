package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/theofalso/timetrack/internal/store"
)

var projectFlag string

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Show tracked time report",
	Long:  `Displays a summary of total tracked time grouped by project, with optional filtering.`,
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

		projectTotals := make(map[string]time.Duration)

		for _, s := range sessions {
			if projectFlag != "" && s.Project != projectFlag {
				continue
			}
			projectTotals[s.Project] += s.Duration()
		}

		if len(projectTotals) == 0 {
			fmt.Printf("No sessions found for project '%s'.\n", projectFlag)
			return nil
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

	reportCmd.Flags().StringVarP(&projectFlag, "project", "p", "", "Filter report by project name")

	rootCmd.AddCommand(reportCmd)
}
