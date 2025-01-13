package userservice

import (
	. "Test.go/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
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
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	result := r.db.Model(&User{}).Where("ID = ?", id).Updates(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	var UpdatedMessage User
	r.db.Model(&User{}).Where("ID = ?", id).First(&UpdatedMessage)
	return UpdatedMessage, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	result := r.db.Where("ID = ?", id).Delete(&User{})
	return result.Error
}
