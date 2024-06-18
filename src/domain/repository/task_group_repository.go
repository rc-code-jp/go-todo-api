package repository

import (
	"context"
	"go-todo-api/src/domain/model"
)

type TaskGroupRepository interface {
	Create(ctx context.Context, m *model.TaskGroup) (*model.TaskGroup, error)
	AllByUserId(ctx context.Context, userId int) (*[]model.TaskGroup, error)
	FindByID(ctx context.Context, id int) (*model.TaskGroup, error)
	Update(ctx context.Context, m *model.TaskGroup) (*model.TaskGroup, error)
	Delete(ctx context.Context, id int) error
}
