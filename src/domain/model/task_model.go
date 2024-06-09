package model

import (
	"time"
)

type Task struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	TaskGroupID  int       `json:"task_group_id"`
	ParentTaskID int       `json:"parent_task_id"`
	Title        string    `json:"title"`
	Date         string    `json:"date"`
	Time         string    `json:"time"`
	Note         string    `json:"note"`
	CompletedAt  time.Time `json:"completed_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

