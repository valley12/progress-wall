package models

import (
	"time"

	"gorm.io/gorm"
)

// UserPermission 用户权限表
type UserPermission struct {
	ID               uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           uint           `json:"user_id" gorm:"not null;index"`
	PermissionString string         `json:"permission_string" gorm:"not null;comment:'权限字符串，;分隔，如*.user.read;project.myproject.task.create'"`
	GrantedBy        uint           `json:"granted_by" gorm:"not null;comment:'授权人ID'"`
	ExpiresAt        *time.Time     `json:"expires_at" gorm:"comment:'权限过期时间，为空表示永不过期'"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	User          User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	GrantedByUser User `json:"granted_by_user,omitempty" gorm:"foreignKey:GrantedBy"`
}

// UserPermissionGroup 用户权限组表
type UserPermissionGroup struct {
	ID                uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupName         string         `json:"group_name" gorm:"size:100;not null;comment:'权限组名称'"`
	PermissionStrings string         `json:"permission_strings" gorm:"type:text;not null;comment:'权限字符串列表，;分隔'"`
	CreatedBy         uint           `json:"created_by" gorm:"not null;comment:'创建人ID'"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Users         []User `json:"users,omitempty" gorm:"many2many:user_permission_group_users"`
	CreatedByUser User   `json:"created_by_user,omitempty" gorm:"foreignKey:CreatedBy"`
}

// UserPermissionGroupUser 用户权限组用户关联表
type UserPermissionGroupUser struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID   uint           `json:"group_id" gorm:"not null;index"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
