package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var task string = "Task"

type requestBody struct {
	Task string `json:"task"`
}

func getTask(c echo.Context) error {
	return c.String(http.StatusOK, "Hello "+task)
}

func postTask(c echo.Context) error {
	var req requestBody

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	task = req.Task

	return c.JSON(http.StatusOK, map[string]string{"message": "Task update successfully", "task": task})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/", getTask)
	e.POST("/task", postTask)

	// запуск сервера на порту 8080
	e.Logger.Fatal(e.Start(":8080"))
}
