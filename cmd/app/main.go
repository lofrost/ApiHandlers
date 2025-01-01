package main

import (
	"net/http"

	"Test.go/internal/database"
	"Test.go/internal/handlers"
	"Test.go/internal/taskService"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/get", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/patch/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
