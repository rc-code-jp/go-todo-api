package repository

import (
	"context"
	"go-todo-api/src/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, m *model.User) (*model.User, error)
	FetchByEmail(ctx context.Context, email string) (*model.User, error)
	FindByID(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, m *model.User) (*model.User, error)
	Delete(ctx context.Context, id int) error
}
