package cmd

import (
	"fmt"

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
		fmt.Println("list called")
	},
}
