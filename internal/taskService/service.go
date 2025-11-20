package taskService

import (
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task string) (TaskSt, error)
	GetAllTasks() ([]TaskSt, error)
	GetTaskByID(id string) (TaskSt, error)
	UpdateTask(id, task string) (TaskSt, error)
	DeleteTask(id string) error
}

type tasService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &tasService{repo: r}
}

// CreateTask implements TaskService.
func (s *tasService) CreateTask(task string) (TaskSt, error) {

	tas := TaskSt{
		ID:   uuid.NewString(),
		Task: task,
	}

	if err := s.repo.CreateTask(tas); err != nil {
		return TaskSt{}, err
	}

	return tas, nil
}

// GetAllTasks implements TaskService.
func (s *tasService) GetAllTasks() ([]TaskSt, error) {
	return s.repo.GetAllTasks()
}

// GetTaskByID implements TaskService.
func (s *tasService) GetTaskByID(id string) (TaskSt, error) {
	return s.repo.GetTaskByID(id)
}

// UpdateTask implements TaskService.
func (s *tasService) UpdateTask(id, task string) (TaskSt, error) {
	tas, err := s.repo.GetTaskByID(id)
	if err != nil {
		return TaskSt{}, err
	}

	tas.Task = task

	if err := s.repo.UpdateTask(tas); err != nil {
		return TaskSt{}, err
	}

	return tas, nil
}

// DeleteTask implements TaskService.
func (s *tasService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
