package main

import (
	"fmt"
	"go-todo-cli/cmd"
	"os"
)

func main() {
	// Check if a subcommand is provided
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'list', or 'complete' subcommands")
		os.Exit(1)
	}

	// Delegate to the cmd package to handle commands
	cmd.Execute(os.Args[1:])
}
