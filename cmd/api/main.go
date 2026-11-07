package main

import (
	"fmt"

	"github.com/matheuscaet/go-simple-ddd/internal/task/domain"
)

func main() {
	// Your code here
	fmt.Println("Hello, World!")
	task := domain.NewTask("1", "Learn Go")
	fmt.Printf("Task created: %+v\n", task)
}
