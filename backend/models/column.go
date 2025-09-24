package models

import (
	"time"

	"gorm.io/gorm"
)

// Column 列表
type Column struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"size:50;not null"`
	Description string         `json:"description" gorm:"size:255"`
	Color       string         `json:"color" gorm:"size:7;default:'#95a5a6';comment:'列颜色，十六进制格式'"`
	Position    int            `json:"position" gorm:"not null;comment:'列在看板中的排序位置'"`
	BoardID     uint           `json:"board_id" gorm:"not null;index"`
	Status      ColumnStatus   `json:"status" gorm:"type:tinyint;default:1;comment:'列状态:1=正常,2=禁用'"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Board Board  `json:"board,omitempty" gorm:"foreignKey:BoardID"`
	Tasks []Task `json:"tasks,omitempty" gorm:"foreignKey:ColumnID"`
}

// ColumnStatus 列状态枚举
type ColumnStatus int

const (
	ColumnStatusActive   ColumnStatus = 1 // 正常
	ColumnStatusDisabled ColumnStatus = 2 // 禁用
)
