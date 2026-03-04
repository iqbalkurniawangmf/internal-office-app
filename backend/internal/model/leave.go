package model

import (
	"time"

	"gorm.io/gorm"
)

type Leave struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	UserID     uint           `json:"user_id" gorm:"not null"`
	StartDate  time.Time      `json:"start_date" gorm:"not null"`
	EndDate    time.Time      `json:"end_date" gorm:"not null"`
	Reason     string         `json:"reason" gorm:"type:text"`
	Status     string         `json:"status" gorm:"type:varchar(50);default:pending"`
	ApprovedBy *uint          `json:"approved_by"` // nullable
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	User        User  `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Approver    *User `json:"approver,omitempty" gorm:"foreignKey:ApprovedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}