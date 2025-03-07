package main

import (
	"SecondProject/internal/database"
	"SecondProject/internal/handlers"
	"SecondProject/internal/taskService"
	"SecondProject/internal/userService"
	"SecondProject/internal/web/tasks"
	"SecondProject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {

	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.TaskNewHandler(tasksService)

	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewService(usersRepo)
	usersHandler := handlers.UserNewHandler(usersService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
