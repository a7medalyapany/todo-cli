package main

import (
	"fmt"
	"strconv"
	"todo-cli/task"
	"todo-cli/utils"
)

func main() {
	// Load tasks from file on startup
	err := task.LoadFromFile()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		choice := utils.GetInput("Enter choice: ")

		switch choice {
		case "1":
			taskName := utils.GetInput("Enter task: ")
			task.AddTask(taskName)
			task.SaveToFile() // Save to file after each operation

		case "2":
			task.ListTasks()

		case "3":
			idStr := utils.GetInput("Enter task ID to mark as done: ")
			taskID, _ := strconv.Atoi(idStr)
			task.MarkDone(taskID)
			task.SaveToFile()

		case "4":
			idStr := utils.GetInput("Enter task ID to delete: ")
			taskID, _ := strconv.Atoi(idStr)
			task.DeleteTask(taskID)
			task.SaveToFile()

		case "5":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
