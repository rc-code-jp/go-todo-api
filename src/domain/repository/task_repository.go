package repository

import (
	"context"
	"go-todo-api/src/domain/model"
)

type TaskRepository interface {
	Create(ctx context.Context, user *model.Task) (*model.Task, error)
	FindByID(ctx context.Context, id int) (*model.Task, error)
	Update(ctx context.Context, user *model.Task) (*model.Task, error)
	Delete(ctx context.Context, id int) error
}
