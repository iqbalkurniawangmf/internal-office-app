package model

import (
	"time"

	"gorm.io/gorm"
)

type Announcement struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"type:varchar(150);not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	CreatedBy uint           `json:"created_by" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relation
	Creator User `json:"creator" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}