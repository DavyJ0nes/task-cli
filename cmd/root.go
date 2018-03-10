package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the entry to the CLI tool
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI todo tool",
}

// Execute runs the RootCmd and outputs the correct os Exit Code
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
