package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Глобальная переменная task
var task string = "world"

// Структура для парсинга JSON из запроса
type RequestBody struct {
	Task string `json:"task"`
}

func main() {
	e := echo.New()

	// GET handler - возвращает "hello, {task}"
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, "+task)
	})

	// POST handler - принимает JSON и обновляет task
	e.POST("/task", func(c echo.Context) error {
		var requestBody RequestBody

		// Парсим JSON из тела запроса
		if err := c.Bind(&requestBody); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid JSON",
			})
		}

		// Обновляем глобальную переменную task
		task = requestBody.Task

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Task updated successfully",
			"task":    task,
		})
	})

	// Запускаем сервер на порту 8080
	e.Logger.Fatal(e.Start(":8080"))
}
