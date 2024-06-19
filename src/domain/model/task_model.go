package model

import (
	"time"
)

type Task struct {
	ID           int       `json:"id" gorm:"primary_key"`
	UserID       int       `json:"user_id"`
	TaskGroupID  int       `json:"task_group_id"`
	ParentTaskID int       `json:"parent_task_id"`
	Title        string    `json:"title" gorm:"type:varchar(256);not null;default:''"`
	Date         string    `json:"date" gorm:"type:date;default:NULL"`
	Time         string    `json:"time" gorm:"type:time;default:NULL"`
	Note         string    `json:"note" gorm:"default:''"`
	CompletedAt  *time.Time `json:"completed_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null;default:current_timestamp"`
}
