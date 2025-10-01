package models

import (
	"time"

	"gorm.io/gorm"
)

// Attachment 附件表
type Attachment struct {
	ID           uint             `json:"id" gorm:"primaryKey;autoIncrement"`
	Filename     string           `json:"filename" gorm:"size:255;not null"`
	OriginalName string           `json:"original_name" gorm:"size:255;not null"`
	FilePath     string           `json:"file_path" gorm:"size:500;not null"`
	FileSize     int64            `json:"file_size" gorm:"not null;comment:'文件大小(字节)'"`
	MimeType     string           `json:"mime_type" gorm:"size:100;not null"`
	TaskID       uint             `json:"task_id" gorm:"not null;index"`
	UploaderID   uint             `json:"uploader_id" gorm:"not null;index"`
	Status       AttachmentStatus `json:"status" gorm:"type:tinyint;default:1;comment:'附件状态:1=正常,2=删除'"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
	DeletedAt    gorm.DeletedAt   `json:"-" gorm:"index"`

	// 关联关系
	Task     Task `json:"task,omitempty" gorm:"foreignKey:TaskID"`
	Uploader User `json:"uploader,omitempty" gorm:"foreignKey:UploaderID"`
}

// AttachmentStatus 附件状态枚举
type AttachmentStatus int

const (
	AttachmentStatusNormal  AttachmentStatus = 1 // 正常
	AttachmentStatusDeleted AttachmentStatus = 2 // 删除
)
