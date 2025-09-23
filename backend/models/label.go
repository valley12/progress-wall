package models

import (
	"time"

	"gorm.io/gorm"
)

// Label 标签表
type Label struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"size:50;not null"`
	Color     string         `json:"color" gorm:"size:7;default:'#3498db';comment:'标签颜色，十六进制格式'"`
	ProjectID uint           `json:"project_id" gorm:"not null;index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Project Project `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
	Tasks   []Task  `json:"tasks,omitempty" gorm:"many2many:task_labels"`
}

// TaskLabel 任务标签关联表
type TaskLabel struct {
	ID      uint `json:"id" gorm:"primaryKey;autoIncrement"`
	TaskID  uint `json:"task_id" gorm:"not null;index"`
	LabelID uint `json:"label_id" gorm:"not null;index"`

	// 关联关系
	Task  Task  `json:"task,omitempty" gorm:"foreignKey:TaskID"`
	Label Label `json:"label,omitempty" gorm:"foreignKey:LabelID"`
}
