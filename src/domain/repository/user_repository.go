// Package repository Domain Service層のリポジトリ
package repository

import (
	"context"
	"go-todo-api/src/domain/model"
)

// UserRepository usersテーブルへのアクセスを行うインターフェース。
type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	FetchByEmail(ctx context.Context, email string) (*model.User, error)
	FetchByID(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id int) error
}
