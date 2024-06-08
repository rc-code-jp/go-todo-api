// Package repository Domain Service層のリポジトリ
package repository

import (
	"go-todo-api/src/domain/model"
)

// UserRepository usersテーブルへのアクセスを行うインターフェース。
type UserRepository interface {
	Create(user *model.User) error
	FetchByEmail(email string) (*model.User, error)
	FetchByID(id int) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
}
