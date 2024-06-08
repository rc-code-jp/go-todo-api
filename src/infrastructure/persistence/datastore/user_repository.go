// Package datastore Infra層のリポジトリ
package datastore

import (
	"context"
	"go-todo-api/src/domain/model"
	"go-todo-api/src/domain/repository"

	"github.com/jinzhu/gorm"
)

// userRepository 構造体
type userRepository struct {
	Conn *gorm.DB
}

// NewUserRepository UserRepositoryを生成する。
func NewUserRepository(Conn *gorm.DB) repository.UserRepository {
	return &userRepository{Conn}
}

// Create 登録
func (repository *userRepository) Create(ctx context.Context, u *model.User) (*model.User, error) {
	err := repository.Conn.Create(u).Error
	return u, err
}

// FetchByEmail メールアドレスが一致するUserを1件取得。
func (repository *userRepository) FetchByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := repository.Conn.Where("email = ?", email).First(&user).Error
	return &user, err
}

// FetchByID IDが一致するUserを1件取得。
func (repository *userRepository) FetchByID(ctx context.Context, id int) (*model.User, error) {
	u := model.User{ID: id}
	if err := repository.Conn.First(&u).Error; err != nil {
		return nil, err
	}
	u.HashedPassword = ""

	return &u, nil
}

// Update 更新
func (repository *userRepository) Update(ctx context.Context, u *model.User) (*model.User, error) {
	err := repository.Conn.Model(u).Update(u).Error
	return u, err
}

// Delete 削除
func (repository *userRepository) Delete(ctx context.Context, id int) error {
	user := model.User{ID: id}
	return repository.Conn.Delete(&user).Error
}
