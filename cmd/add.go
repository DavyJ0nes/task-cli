package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/davyj0nes/task-cli/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a command to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		// join all space seperated words after add command into a single string
		taskName := strings.Join(args, " ")

		if err := db.CreateTask(taskName); err != nil {
			fmt.Printf("Error Adding Task:\n  %v\n", err)
			os.Exit(1)
		}
	},
}
