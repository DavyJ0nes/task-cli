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
		if err := db.ListTasks(); err != nil {
			fmt.Printf("Error Listing Tasks:\n  %v\n", err)
			os.Exit(1)
		}
	},
}
