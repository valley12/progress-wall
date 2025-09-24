package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户表
// 存储系统用户的基本信息和状态
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement" comment:"用户唯一标识符，自增主键"`
	Username  string         `json:"username" gorm:"uniqueIndex;size:50;not null" comment:"用户名，唯一标识，最大50字符"`
	Email     string         `json:"email" gorm:"uniqueIndex;size:100;not null" comment:"邮箱地址，唯一标识，最大100字符"`
	Password  string         `json:"-" gorm:"size:255;not null" comment:"密码哈希值，不在JSON中返回，最大255字符"`
	Nickname  string         `json:"nickname" gorm:"size:50" comment:"用户昵称，用于显示，最大50字符"`
	Avatar    string         `json:"avatar" gorm:"size:255" comment:"头像URL地址，最大255字符"`
	Phone     string         `json:"phone" gorm:"size:20" comment:"手机号码，最大20字符"`
	Status    UserStatus     `json:"status" gorm:"type:tinyint;default:1;comment:'用户状态:1=正常,2=禁用,3=删除'" comment:"用户状态，1=正常，2=禁用，3=删除"`
	LastLogin *time.Time     `json:"last_login" comment:"最后登录时间，可为空"`
	CreatedAt time.Time      `json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time      `json:"updated_at" comment:"更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" comment:"软删除时间，用于软删除标记"`

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
	UserStatusEnabled  UserStatus = 1 // 正常
	UserStatusDisabled UserStatus = 2 // 禁用
	UserStatusDeleted  UserStatus = 3 // 删除
)
