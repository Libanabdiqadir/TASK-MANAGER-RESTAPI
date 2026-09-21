package main

import (
	"fmt"
	"net/http"
)


func main() {
	taskManager := &TaskManager{}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /tasks", taskManager.CreateTaskHandler)
	mux.HandleFunc("GET /tasks", taskManager.GetAllTasksHandler)

	mux.HandleFunc("GET /tasks/{id}", taskManager.GetTaskByIDHandler)
	mux.HandleFunc("PUT /tasks/{id}", taskManager.UpdateTaskHandler)
	mux.HandleFunc("DELETE /tasks/{id}", taskManager.DeleteTaskHandler)

	handler := loggingMiddleware(mux)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}