package model

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"type:varchar(100);not null;unique"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relation
	Users []User `json:"users,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}