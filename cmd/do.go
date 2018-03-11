package cmd

import (
	"fmt"
	"os"

	"github.com/davyj0nes/task-cli/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(doCmd)
}

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Completes a task",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.CompleteTask(args); err != nil {
			fmt.Printf("Error Completing Task:\n  %v\n", err)
			os.Exit(1)
		}
	},
}
