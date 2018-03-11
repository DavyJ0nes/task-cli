package cmd

import (
	"fmt"
	"os"

	"github.com/davyj0nes/task-cli/db"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the current open tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listTasks(); err != nil {
			fmt.Printf("Error Listing Tasks:\n  %v\n", err)
			os.Exit(1)
		}
	},
}

func listTasks() error {
	tasks, err := db.AllTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks, you're clear")
		return nil
	}

	for i, task := range tasks {
		fmt.Fprintf(os.Stdout, "%d. | %s\n", i+1, task.Name)
	}
	return nil
}
