package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	DepartmentID *uint          `json:"department_id"` // nullable
	Name         string         `json:"name" gorm:"type:varchar(100);not null"`
	Email        string         `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password     string         `json:"-" gorm:"type:varchar(255);not null"`
	Role         string         `json:"role" gorm:"type:varchar(50);default:employee"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	Department *Department `json:"department,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}