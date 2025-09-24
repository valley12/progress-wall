package models

import (
	"time"

	"gorm.io/gorm"
)

// Comment 评论表
type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	TaskID    uint           `json:"task_id" gorm:"not null;index"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	ParentID  *uint          `json:"parent_id" gorm:"index;comment:'父评论ID，用于回复评论'"`
	Status    CommentStatus  `json:"status" gorm:"type:tinyint;default:1;comment:'评论状态:1=正常,2=隐藏,3=删除'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Task    Task      `json:"task,omitempty" gorm:"foreignKey:TaskID"`
	User    User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Parent  *Comment  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies []Comment `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
}

// CommentStatus 评论状态枚举
type CommentStatus int

const (
	CommentStatusNormal  CommentStatus = 1 // 正常
	CommentStatusHidden  CommentStatus = 2 // 隐藏
	CommentStatusDeleted CommentStatus = 3 // 删除
)
