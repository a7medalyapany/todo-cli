package task

import "fmt"

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var Tasks []Task
var idCounter int

func AddTask(taskName string) {
	idCounter++
	task := Task{ID: idCounter, Name: taskName, Done: false}
	Tasks = append(Tasks, task)
	fmt.Println("Task added!")
}

func ListTasks() {
	if len(Tasks) == 0 {
		fmt.Println("No tasks!")
		return
	}

	for _, task := range Tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d - %s [%s]\n", task.ID, task.Name, status)
	}
}

func MarkDone(taskID int) {
	for i, task := range Tasks {
		if task.ID == taskID {
			Tasks[i].Done = true
			fmt.Println("Task marked as done!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func DeleteTask(taskID int) {
	for i, task := range Tasks {
		if task.ID == taskID {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			fmt.Println("Task deleted!")
			return
		}
	}
	fmt.Println("Task not found!")
}
