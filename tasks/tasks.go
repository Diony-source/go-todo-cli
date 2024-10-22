package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// File where tasks are stored
const taskFile = "tasks.txt"

// AddTask adds a new task to the list
func AddTask(task string) {
	f, err := os.OpenFile(taskFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString(task + "\n"); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Task added: %s\n", task)
}

// ListTasks lists all tasks from the file
func ListTasks() {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Your tasks:")
	i := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", i, scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading tasks:", err)
		os.Exit(1)
	}
}

// CompleteTask marks a task as complete (to be implemented)
func CompleteTask(taskNumber string) {
	//
