package model

import (
	"time"
)

type User struct {
	ID             int        `json:"id" gorm:"primary_key"`
	Name           string     `json:"name" gorm:"type:varchar(256);not null;default:''"`
	Email          string     `json:"email" gorm:"type:varchar(256);not null;default:'';unique"`
	HashedPassword string     `json:"hashed_password" gorm:"type:varchar(256);not null;default:''"`
	ImageFilePath  string     `json:"image_file_path" gorm:"type:varchar(256);not null;default:''"`
	DeletedAt      *time.Time `json:"deleted_at"`
	CreatedAt      time.Time  `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"not null;default:current_timestamp"`
}
