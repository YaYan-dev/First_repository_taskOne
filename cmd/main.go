package main

import (
	"FIRST_REPOSITORY_TASKONE/internal/db"
	"FIRST_REPOSITORY_TASKONE/internal/handlers"
	"FIRST_REPOSITORY_TASKONE/internal/taskService"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	tasRepo := taskService.NewTaskRepository(database)
	tasService := taskService.NewTaskService(tasRepo)
	tasHandler := handlers.NewTaskHandler(tasService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", tasHandler.GetTask)
	e.POST("/task", tasHandler.PostTask)
	e.PATCH("/task/:id", tasHandler.PatchTask)
	e.DELETE("task/:id", tasHandler.DeleteTask)

	// запуск сервера на порту 8080
	e.Logger.Fatal(e.Start(":8080"))
}
