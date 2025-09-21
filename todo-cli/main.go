package main

import (
	"fmt"
)

func main() {
	todoManage := NewManage()
	var command string

	running := true
	for running {
		fmt.Print("Enter your command:")
		_, _ = fmt.Scanln(&command)
		switch command {
		case "list":
			fmt.Println("ID\tTask Name\tCompleted")
			for _, todo := range *todoManage.todoList {
				fmt.Println(fmt.Sprintf("%d\t%s\t%t", todo.ID, todo.Task, todo.Completed))
			}
			break
		case "add":
			fmt.Print("Enter your task:")
			var task string
			_, _ = fmt.Scanln(&task)
			todoManage.Add(task)
			break
		case "complete":
			fmt.Print("Enter your task ID:")
			var taskID int
			_, _ = fmt.Scanln(&taskID)
			todoManage.Complete(taskID)
			break
		case "delete":
			fmt.Print("Enter your task ID:")
			var taskID int
			_, _ = fmt.Scanln(&taskID)
			todoManage.Delete(taskID)
			break
		case "exit":
			running = false
			break
		default:
			fmt.Print("Invalid command")
		}
	}
}
