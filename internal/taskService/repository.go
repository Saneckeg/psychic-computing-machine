package taskService

import (
	"gorm.io/gorm"
	"log"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task map[string]interface{}) (Task, error)
	DeleteTaskByID(id uint) (Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var task []Task

	err := r.db.Find(&task).Error
	if err != nil {
		log.Println("Ошибка при получении данных:", err)
	}

	return task, err
}

func (r *taskRepository) UpdateTaskByID(id uint, updates map[string]interface{}) (Task, error) {
	var task Task

	if err := r.db.First(&task, id).Error; err != nil {
		return task, err
	}

	result := r.db.Model(&task).Updates(updates)
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return task, err
	}

	if err := r.db.Delete(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task map[string]interface{}) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) (Task, error) {
	return s.repo.DeleteTaskByID(id)
}
