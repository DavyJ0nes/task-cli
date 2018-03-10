package cmd

import (
	"fmt"
	"os"
	"strings"

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
		if err := addTask(args); err != nil {
			fmt.Printf("Error Adding Task:\n  %v\n", err)
			os.Exit(1)
		}
	},
}

// addTask takes the arguments after the add command and creates a new entry in the database
func addTask(args []string) error {
	taskName := strings.Join(args, " ")
	fmt.Printf("adding task: '%s'\n", taskName)
	return nil
}
