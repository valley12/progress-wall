package models

import (
	"time"

	"gorm.io/gorm"
)

// Board 看板表
type Board struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"size:100;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Color       string         `json:"color" gorm:"size:7;default:'#3498db';comment:'看板颜色，十六进制格式'"`
	Status      BoardStatus    `json:"status" gorm:"type:tinyint;default:1;comment:'看板状态:1=活跃,2=归档'"`
	ProjectID   uint           `json:"project_id" gorm:"not null;index"`
	OwnerID     uint           `json:"owner_id" gorm:"not null;index"`
	Position    int            `json:"position" gorm:"default:0;comment:'看板在项目中的排序位置'"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Project Project  `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
	Owner   User     `json:"owner,omitempty" gorm:"foreignKey:OwnerID"`
	Columns []Column `json:"columns,omitempty" gorm:"foreignKey:BoardID"`
}

// BoardStatus 看板状态枚举
type BoardStatus int

const (
	BoardStatusActive   BoardStatus = 1 // 活跃
	BoardStatusArchived BoardStatus = 2 // 归档
)
