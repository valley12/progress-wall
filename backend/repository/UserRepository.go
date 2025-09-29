package repository

import (
	"progress-wall-backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// 保存用户
func (r *userRepository) Save(user *models.User) (error) {
	return r.db.Create(user).Error
}

// 根据用户名查找用户
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 根据邮箱查找用户
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
