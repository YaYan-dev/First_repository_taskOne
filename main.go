package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TaskSt struct {
	ID     string `json:"id"`
	Build  string `json:"build"`
	Result string `json:"result"`
}

type RequestBody struct {
	Build string `json:"build"`
}

var task = []TaskSt{}

func buildTask(build string) string {
	return "Hello, " + build
}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, task)
}

func postTask(c echo.Context) error {
	var req RequestBody

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	result := buildTask(req.Build)

	tas := TaskSt{
		ID:     uuid.NewString(),
		Build:  req.Build,
		Result: result,
	}
	task = append(task, tas)
	return c.JSON(http.StatusCreated, tas)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")

	var req RequestBody

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	result := buildTask(req.Build)

	for i, taskItem := range task {
		if taskItem.ID == id {
			task[i].Build = req.Build
			task[i].Result = result
			return c.JSON(http.StatusOK, task[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task (id) not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	for i, taskItem := range task {
		if taskItem.ID == id {
			task = append(task[:i], task[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task (id) not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", getTask)
	e.POST("/task", postTask)
	e.PATCH("/task/:id", patchTask)
	e.DELETE("task/:id", deleteTask)

	// запуск сервера на порту 8080
	e.Logger.Fatal(e.Start(":8080"))
}
