package main

import (
	"log"

	"Test.go/internal/database"
	"Test.go/internal/handlers"
	"Test.go/internal/taskService"
	"Test.go/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatal("Failed to start with err:", err)
	}

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)

	tasks.RegisterHandlers(e, strictHandler)

	if err = e.Start(":8080"); err != nil {
		log.Fatal("Failed to start with err:", err)
	}

	//router := mux.NewRouter()
	//router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	//router.HandleFunc("/api/get", handler.GetTaskHandler).Methods("GET")
	//router.HandleFunc("/api/patch/{id}", handler.PatchTaskHandler).Methods("PATCH")
	//router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	//http.ListenAndServe(":8080", router)

}
