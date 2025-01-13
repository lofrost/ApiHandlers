package main

import (
	"log"

	"Test.go/internal/database"
	"Test.go/internal/handlers"
	"Test.go/internal/models"
	"Test.go/internal/taskService"
	userservice "Test.go/internal/userService"
	"Test.go/internal/web/tasks"
	"Test.go/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to start with err:", err)
	}

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	userRepo := userservice.NewUserRepository(database.DB)
	userService := userservice.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)
	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err = e.Start(":8080"); err != nil {
		log.Fatal("Failed to start with err:", err)
	}

}
