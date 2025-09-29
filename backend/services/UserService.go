package services

import (
	"errors"
	"progress-wall-backend/models"
	"progress-wall-backend/repository"
)

type UserService interface {
	Register(username, email, password string) (*models.User, error)
	Login(email, password string) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// Register 注册新用户
func (s *userService) Register(username, email, password string) (*models.User, error) {
	// 检查用户名或邮箱是否已存在
	if user, err := s.userRepo.FindByUsername(username); err == nil {
		return user, errors.New("username already exists")
	}
	if user, err := s.userRepo.FindByEmail(email); err == nil {
		return user, errors.New("email already exists")
	}

	//创建用户实例
	user := &models.User{
		Username: username,
		Email:    email,
	}

	//加密密码
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	return user, s.userRepo.Save(user)
}

// Login 用户登录
func (s *userService) Login(email, password string) (*models.User, error) {
	// 根据邮箱查找用户
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// 检查密码
	if !user.CheckPassword(password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
