package taskService

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTasksByUserID(userID int) ([]Task, error)
	UpdateTaskByID(id uint, task interface{}) (Task, error)
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
		log.Printf("Ошибка при создании задачи: %v", result.Error)
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

func (r *taskRepository) GetTasksByUserID(userID int) ([]Task, error) {
	var tasks []Task

	err := r.db.Where("user_id =?", userID).Find(&tasks).Error
	if err != nil {
		log.Println("Ошибка при получении задач пользователя:", err)
	}

	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, updates interface{}) (Task, error) {
	var task Task

	if err := r.db.First(&task, id).Error; err != nil {
		return task, err
	}

	var updatesMap map[string]interface{}

	// Если updates уже map[string]interface{}, просто используем его
	if casted, ok := updates.(map[string]interface{}); ok {
		updatesMap = casted
	} else {
		// Если это структура, конвертируем её в map[string]interface{}
		bytes, err := json.Marshal(updates)
		if err != nil {
			return task, err
		}
		if err := json.Unmarshal(bytes, &updatesMap); err != nil {
			return task, err
		}
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
