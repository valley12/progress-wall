package models

import (
	"time"

	"gorm.io/gorm"
)

// Project 项目表
type Project struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"size:100;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Status      ProjectStatus  `json:"status" gorm:"type:tinyint;default:1;comment:'项目状态:1=进行中,2=已完成,3=已暂停,4=已取消'"`
	StartDate   *time.Time     `json:"start_date"`
	EndDate     *time.Time     `json:"end_date"`
	OwnerID     uint           `json:"owner_id" gorm:"not null;index"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Owner   User    `json:"owner,omitempty" gorm:"foreignKey:OwnerID" comment:"项目所有者，必须有项目根权限*.project.project_id.*"`
	Boards  []Board `json:"boards,omitempty" gorm:"foreignKey:ProjectID"`
	Members []User  `json:"members,omitempty" gorm:"many2many:project_members"`
}

// ProjectMember 项目成员关联表
type ProjectMember struct {
	ID                uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectID         uint           `json:"project_id" gorm:"not null;index"`
	UserID            uint           `json:"user_id" gorm:"not null;index"`
	PorjectPermission string         `json:"project_permission" gorm:"type:text;comment:'项目权限，如*.project.project_id.*'"`
	JoinedAt          time.Time      `json:"joined_at"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Project Project `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
	User    User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// ProjectStatus 项目状态枚举
type ProjectStatus int

const (
	ProjectStatusActive    ProjectStatus = 1 // 进行中
	ProjectStatusCompleted ProjectStatus = 2 // 已完成
	ProjectStatusPaused    ProjectStatus = 3 // 已暂停
	ProjectStatusCancelled ProjectStatus = 4 // 已取消
)

// VisibilityStatus 项目可见性枚举
type ProjectVisibility int

const (
	ProjectVisibilityPublic  ProjectVisibility = 1 // 公开
	ProjectVisibilityPrivate ProjectVisibility = 2 // 私有
)
