package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/davyj0nes/task-cli/db"

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

func completeTask(args []string) error {
	var ids []int
	// convert arguments into integers
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return errors.Errorf("Could not parse Argument: %s", arg)
		}
		ids = append(ids, id)
	}

	tasks, err := db.AllTasks()
	if err != nil {
		return errors.Wrap(err, "Could not get all tasks")
	}
	// delete tasks from database
	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Println("Invalid task number:", id)
			continue
		}
		task := tasks[id-1]
		if err := db.DeleteTask(task.ID); err != nil {
			return errors.Wrapf(err, "Could not delete task: %s", task.Name)
		}
		// print out for user
		fmt.Printf("Marked %s as Completed\n", task.Name)
	}

	return nil
}
