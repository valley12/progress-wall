package database

import (
	"log"
	"progress-wall-backend/models"

	"gorm.io/gorm"
)

// Seed 初始化基础数据
func Seed(db *gorm.DB) error {
	log.Println("开始初始化基础数据...")

	// 创建默认用户
	if err := createDefaultUsers(db); err != nil {
		return err
	}

	// 创建默认权限
	if err := createDefaultPermissions(db); err != nil {
		return err
	}

	log.Println("基础数据初始化完成")
	return nil
}

func createDefaultUsers(_ *gorm.DB) error {
	// TODO: 创建默认用户，目前密码加密方法还没确定
	return nil
}

func createDefaultPermissions(db *gorm.DB) error {
	adminPermissionGroup := models.UserPermissionGroup{
		GroupName:         "admin",
		PermissionStrings: "*",
	}
	db.Create(&adminPermissionGroup)
	guestPermissionGroup := models.UserPermissionGroup{
		GroupName:         "guest",
		PermissionStrings: "users.list;users.detail",
	}
	db.Create(&guestPermissionGroup)
	return nil
}
