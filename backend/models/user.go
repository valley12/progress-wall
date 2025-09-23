package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户表
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string         `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Email     string         `json:"email" gorm:"uniqueIndex;size:100;not null"`
	Password  string         `json:"-" gorm:"size:255;not null"`
	Nickname  string         `json:"nickname" gorm:"size:50"`
	Avatar    string         `json:"avatar" gorm:"size:255"`
	Phone     string         `json:"phone" gorm:"size:20"`
	Status    UserStatus     `json:"status" gorm:"type:tinyint;default:1;comment:'用户状态:1=正常,2=禁用,3=删除'"`
	LastLogin *time.Time     `json:"last_login"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Permissions      []UserPermission      `json:"permissions,omitempty" gorm:"foreignKey:UserID"`
	PermissionGroups []UserPermissionGroup `json:"permission_groups,omitempty" gorm:"many2many:user_permission_group_users"`
	Projects         []Project             `json:"projects,omitempty" gorm:"many2many:project_members"`
	CreatedTasks     []Task                `json:"created_tasks,omitempty" gorm:"foreignKey:CreatorID"`
	AssignedTasks    []Task                `json:"assigned_tasks,omitempty" gorm:"foreignKey:AssigneeID"`
	Comments         []Comment             `json:"comments,omitempty" gorm:"foreignKey:UserID"`
}

// UserStatus 用户状态枚举
type UserStatus int

const (
	UserStatusActive   UserStatus = 1 // 正常
	UserStatusDisabled UserStatus = 2 // 禁用
	UserStatusDeleted  UserStatus = 3 // 删除
)
