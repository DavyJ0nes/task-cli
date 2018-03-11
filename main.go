package main

import (
	"github.com/davyj0nes/task-cli/cmd"
	"github.com/davyj0nes/task-cli/db"
)

func main() {
	db.Initialise("tasks.db")
	cmd.Execute()
}
