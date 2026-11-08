package main

import (
	"fmt"

	"github.com/matheuscaet/go-simple-ddd/internal/task/application"
	repoTask "github.com/matheuscaet/go-simple-ddd/internal/task/infrastructure/repository"
)

type TaskHandler struct {
	service *application.TaskService
}

func main() {
	// Your code here
	fmt.Println("Hello, World!")
	handler := &TaskHandler{
		service: application.NewTaskService(repoTask.NewInMemoryTaskRepository()),
	}
	task, err := handler.service.Create(&application.TaskInput{Title: "My first task"})
	if err != nil {
		fmt.Printf("Error creating task: %v\n", err)
		return
	}
	fmt.Printf("Task created: %+v\n", task)
}
