package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:255;not null"`
	Email    string `gorm:"uniqueIndex;size:255;not null"`
	Password string `gorm:"size:255;not null"`
}

// SetPassword 加密密码
func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// checkPassword 检查密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	//如果err = nil 说明密码匹配
	return err == nil
}
