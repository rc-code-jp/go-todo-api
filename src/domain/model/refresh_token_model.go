package model

import (
	"time"
)

type RefreshToken struct {
	ID          int        `json:"id" gorm:"primary_key"`
	UUID        string     `json:"uuid" gorm:"type:varchar(255);not null"`
	HashedToken string     `json:"hashed_token" gorm:"type:varchar(255);not null"`
	UserID      int        `json:"user_id" gorm:"not null"`
	Revoked     bool       `json:"revoked" gorm:"not null;default:0"`
	CreatedAt   time.Time  `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"not null;default:current_timestamp"`
}

