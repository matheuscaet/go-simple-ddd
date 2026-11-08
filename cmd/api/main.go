package main

import (
	"log"
	"net/http"

	taskApp "github.com/matheuscaet/go-simple-ddd/internal/task/application"
	taskInfra "github.com/matheuscaet/go-simple-ddd/internal/task/infra"
	taskHttp "github.com/matheuscaet/go-simple-ddd/internal/task/interfaces/http"
)

func main() {
	address := "localhost:8090"

	taskService := taskApp.NewTaskService(taskInfra.NewMemoryRepository())

	mux := http.NewServeMux()
	taskHandler := taskHttp.NewHandler(taskService)
	taskHandler.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	log.Printf("Server is running on %s", address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
