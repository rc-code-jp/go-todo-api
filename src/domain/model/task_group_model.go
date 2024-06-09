package model

import (
	"time"
)

type TaskGroup struct {
	ID         int        `json:"id" gorm:"primary_key"`
	UserID     int        `json:"user_id" gorm:"not null"`
	Name       string     `json:"name" gorm:"type:varchar(255)"`
	DeletedAt  *time.Time `json:"deleted_at"`
	CreatedAt  time.Time  `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"not null;default:current_timestamp"`
	// Relationship
	Tasks      []Task     `json:"tasks" gorm:"foreignKey:TaskGroupID;references:ID"`
}

