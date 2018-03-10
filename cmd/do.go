package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"

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
		if err := completeTask(args); err != nil {
			fmt.Printf("Error Completing Task:\n  %v\n", err)
			os.Exit(1)
		}
	},
}

// completeTask takes the ID of a task and updates its status to complete
func completeTask(args []string) error {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return errors.Errorf("Could not parse Argument: %s", arg)
		}
		ids = append(ids, id)
	}
	fmt.Println(ids)
	return nil
}
