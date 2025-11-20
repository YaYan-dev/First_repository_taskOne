package taskService

import "gorm.io/gorm"

//Основные методы CRUD - Create, Read, Update, Delete

type TaskRepository interface {
	CreateTask(tas TaskSt) error
	GetAllTasks() ([]TaskSt, error)
	GetTaskByID(id string) (TaskSt, error)
	UpdateTask(tas TaskSt) error
	DeleteTask(id string) error
}

type tasRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &tasRepository{db: db}
}

func (r *tasRepository) CreateTask(tas TaskSt) error {
	return r.db.Create(&tas).Error
}

func (r *tasRepository) GetAllTasks() ([]TaskSt, error) {
	var tasks []TaskSt
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *tasRepository) GetTaskByID(id string) (TaskSt, error) {
	var tas TaskSt
	err := r.db.First(&tas, "id = ?", id).Error
	return tas, err
}

func (r *tasRepository) UpdateTask(tas TaskSt) error {
	return r.db.Save(&tas).Error
}

func (r *tasRepository) DeleteTask(id string) error {
	return r.db.Delete(&TaskSt{}, "id = ?", id).Error
}
