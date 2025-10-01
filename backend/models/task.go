package models

import (
	"time"

	"gorm.io/gorm"
)

// Task 任务表
type Task struct {
	ID             uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title          string         `json:"title" gorm:"size:200;not null"`
	Description    string         `json:"description" gorm:"type:text"`
	Priority       TaskPriority   `json:"priority" gorm:"type:tinyint;default:2;comment:'任务优先级:1=低,2=中,3=高,4=紧急'"`
	Status         TaskStatus     `json:"status" gorm:"type:tinyint;default:1;comment:'任务状态:1=待办,2=进行中,3=已完成,4=已取消'"`
	Position       int            `json:"position" gorm:"not null;comment:'任务在列中的排序位置'"`
	DueDate        *time.Time     `json:"due_date"`
	StartDate      *time.Time     `json:"start_date"`
	EndDate        *time.Time     `json:"end_date"`
	EstimatedHours *float64       `json:"estimated_hours" gorm:"type:decimal(8,2);comment:'预估工时(小时)'"`
	ActualHours    *float64       `json:"actual_hours" gorm:"type:decimal(8,2);comment:'实际工时(小时)'"`
	ColumnID       uint           `json:"column_id" gorm:"not null;index"`
	CreatorID      uint           `json:"creator_id" gorm:"not null;index"`
	AssigneeID     *uint          `json:"assignee_id" gorm:"index"`
	ProjectID      uint           `json:"project_id" gorm:"not null;index"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Column      Column       `json:"column,omitempty" gorm:"foreignKey:ColumnID"`
	Creator     User         `json:"creator,omitempty" gorm:"foreignKey:CreatorID"`
	Assignee    *User        `json:"assignee,omitempty" gorm:"foreignKey:AssigneeID"`
	Project     Project      `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
	Comments    []Comment    `json:"comments,omitempty" gorm:"foreignKey:TaskID"`
	Attachments []Attachment `json:"attachments,omitempty" gorm:"foreignKey:TaskID"`
	Labels      []Label      `json:"labels,omitempty" gorm:"many2many:task_labels"`
}

// TaskPriority 任务优先级枚举
type TaskPriority int

const (
	TaskPriorityLow    TaskPriority = 1 // 低
	TaskPriorityMedium TaskPriority = 2 // 中
	TaskPriorityHigh   TaskPriority = 3 // 高
	TaskPriorityUrgent TaskPriority = 4 // 紧急
)

// TaskStatus 任务状态枚举
type TaskStatus int

const (
	TaskStatusTodo       TaskStatus = 1 // 待办
	TaskStatusInProgress TaskStatus = 2 // 进行中
	TaskStatusCompleted  TaskStatus = 3 // 已完成
	TaskStatusCancelled  TaskStatus = 4 // 已取消
)
