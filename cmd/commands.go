package cmd

import (
	"flag"
	"fmt"
	"go-todo-cli/tasks"
	"os"
)

// Execute processes the CLI commands
func Execute(args []string) {
	// Define subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)

	// Add subcommand flag for task description
	addTask := addCmd.String("task", "", "Task to be added to the to-do list")

	// Switch on the first argument to identify the subcommand
	switch args[0] {
	case "add":
		addCmd.Parse(args[1:])
		if *addTask == "" {
			fmt.Println("Please provide a task with -task flag.")
			os.Exit(1)
		}
		tasks.AddTask(*addTask)

	case "list":
		listCmd.Parse(args[1:])
		tasks.ListTasks()

	case "complete":
		completeCmd.Parse(args[1:])
		if len(args[1:]) < 1 {
			fmt.Println("Please provide the task number to complete.")
			os.Exit(1)
		}
		tasks.CompleteTask(args[1])

	default:
		fmt.Println("Unknown command. Use 'add', 'list', or 'complete'.")
	}
}
