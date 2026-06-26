package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [project]",
	Short: "Start tracking a project",
	Long:  `Starts a new time session for the specified project.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		startTime := time.Now().Format(time.RFC3339)

		fmt.Printf("Starting session for project '%s' at %s\n", projectName, startTime)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
