package userService

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	CreateUser(task User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, task interface{}) (User, error)
	DeleteUserByID(id uint) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var user []User

	err := r.db.Find(&user).Error
	if err != nil {
		log.Println("Ошибка при получении данных:", err)
	}

	return user, err
}

func (r *userRepository) UpdateUserByID(id uint, updates interface{}) (User, error) {
	var user User

	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}

	var updatesMap map[string]interface{}

	// Если updates уже map[string]interface{}, просто используем его
	if casted, ok := updates.(map[string]interface{}); ok {
		updatesMap = casted
	} else {
		// Если это структура, конвертируем её в map[string]interface{}
		bytes, err := json.Marshal(updates)
		if err != nil {
			return user, err
		}
		if err := json.Unmarshal(bytes, &updatesMap); err != nil {
			return user, err
		}
	}

	result := r.db.Model(&user).Updates(updates)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (r *userRepository) DeleteUserByID(id uint) (User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
