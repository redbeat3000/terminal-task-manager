package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const dbFile = "tasks.json"

type Task struct {
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func loadTasks() ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(dbFile)
	if err == nil {
		json.Unmarshal(data, &tasks)
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dbFile, data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  add [task]      → Add a new task")
		fmt.Println("  list            → List all tasks")
		fmt.Println("  done [index]    → Mark a task as complete")
		fmt.Println("  delete [index]  → Delete a task")
		return
	}

	command := os.Args[1]
	tasks, _ := loadTasks()

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Enter task name.")
			return
		}
		task := Task{Name: os.Args[2], Complete: false}
		tasks = append(tasks, task)
		saveTasks(tasks)
		fmt.Println("Task added.")

	case "list":
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for i, t := range tasks {
			status := " "
			if t.Complete {
				status = "✓"
			}
			fmt.Printf("[%s] %d: %s\n", status, i+1, t.Name)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Enter task number.")
			return
		}
		idx, _ := strconv.Atoi(os.Args[2])
		if idx <= 0 || idx > len(tasks) {
			fmt.Println("Invalid task number.")
			return
		}
		tasks[idx-1].Complete = true
		saveTasks(tasks)
		fmt.Println("Task marked as complete.")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Enter task number.")
			return
		}
		idx, _ := strconv.Atoi(os.Args[2])
		if idx <= 0 || idx > len(tasks) {
			fmt.Println("Invalid task number.")
			return
		}
		tasks = append(tasks[:idx-1], tasks[idx:]...)
		saveTasks(tasks)
		fmt.Println("Task deleted.")

	default:
		fmt.Println("Unknown command.")
	}
}
