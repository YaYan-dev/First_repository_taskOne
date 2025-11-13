package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&TaskSt{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

type TaskSt struct {
	ID    string `gorm:"primaryKey" json:"id"`
	Build string `json:"build"`
}

type RequestBody struct {
	Build string `json:"build"`
}

//Основные методы ORM - Create, Find, Update, Delete

func getTask(c echo.Context) error {
	var tasks []TaskSt

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req RequestBody

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	tas := TaskSt{
		ID:    uuid.NewString(),
		Build: req.Build,
	}

	if err := db.Create(&tas).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not add tasks"})
	}

	return c.JSON(http.StatusCreated, tas)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")

	var req RequestBody

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	var tas TaskSt
	if err := db.First(&tas, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not find task"})
	}

	tas.Build = req.Build

	if err := db.Save(&tas).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, tas)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&TaskSt{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()

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
