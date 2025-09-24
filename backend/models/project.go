package models

import (
	"time"

	"gorm.io/gorm"
)

// Project 项目表
// 存储项目的基本信息和状态
type Project struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement" comment:"项目唯一标识符，自增主键"`
	Name        string         `json:"name" gorm:"size:100;not null" comment:"项目名称，最大100字符，必填"`
	Description string         `json:"description" gorm:"type:text" comment:"项目描述，文本类型"`
	Status      ProjectStatus  `json:"status" gorm:"type:tinyint;default:1;comment:'项目状态:1=进行中,2=已完成,3=已暂停,4=已取消'" comment:"项目状态，1=进行中，2=已完成，3=已暂停，4=已取消"`
	StartDate   *time.Time     `json:"start_date" comment:"项目开始时间，可为空"`
	EndDate     *time.Time     `json:"end_date" comment:"项目结束时间，可为空"`
	OwnerID     uint           `json:"owner_id" gorm:"not null;index" comment:"项目所有者用户ID，必填，建立索引"`
	CreatedAt   time.Time      `json:"created_at" comment:"创建时间"`
	UpdatedAt   time.Time      `json:"updated_at" comment:"更新时间"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index" comment:"软删除时间，用于软删除标记"`

	// 关联关系
	Owner   User    `json:"owner,omitempty" gorm:"foreignKey:OwnerID" comment:"项目所有者，必须有项目根权限*.project.project_id.*"`
	Boards  []Board `json:"boards,omitempty" gorm:"foreignKey:ProjectID"`
	Members []User  `json:"members,omitempty" gorm:"many2many:project_members"`
}

// ProjectMember 项目成员关联表
// 存储项目与用户的关联关系及权限信息
type ProjectMember struct {
	ID                uint           `json:"id" gorm:"primaryKey;autoIncrement" comment:"关联记录唯一标识符，自增主键"`
	ProjectID         uint           `json:"project_id" gorm:"not null;index" comment:"项目ID，必填，建立索引"`
	UserID            uint           `json:"user_id" gorm:"not null;index" comment:"用户ID，必填，建立索引"`
	ProjectPermission string         `json:"project_permission" gorm:"type:text;comment:'项目权限，如*.project.project_id.*'" comment:"项目权限字符串，如*.project.project_id.*"`
	JoinedAt          time.Time      `json:"joined_at" comment:"加入项目时间"`
	CreatedAt         time.Time      `json:"created_at" comment:"创建时间"`
	UpdatedAt         time.Time      `json:"updated_at" comment:"更新时间"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index" comment:"软删除时间，用于软删除标记"`

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
