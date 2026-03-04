package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	AssignedTo  uint           `json:"assigned_to" gorm:"not null"`
	CreatedBy   uint           `json:"created_by" gorm:"not null"`
	Title       string         `json:"title" gorm:"type:varchar(150);not null"`
	Description string         `json:"description" gorm:"type:text"`
	Status      string         `json:"status" gorm:"type:varchar(50);default:todo"`
	DueDate     *time.Time     `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	Assignee User `json:"assignee" gorm:"foreignKey:AssignedTo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Creator  User `json:"creator" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}