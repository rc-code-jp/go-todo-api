package datastore

import (
	"context"
	"go-todo-api/src/domain/model"
	"go-todo-api/src/domain/repository"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) repository.UserRepository {
	return &userRepository{Conn}
}

func (repository *userRepository) Create(ctx context.Context, m *model.User) (*model.User, error) {
	err := repository.Conn.Create(m).Error
	return m, err
}

func (repository *userRepository) FetchByEmail(ctx context.Context, email string) (*model.User, error) {
	var m model.User
	err := repository.Conn.Where("email = ?", email).First(&m).Error
	return &m, err
}

func (repository *userRepository) FindByID(ctx context.Context, id int) (*model.User, error) {
	m := model.User{ID: id}
	err := repository.Conn.First(&m).Error;
	if err != nil {
		return nil, err
	}
	m.HashedPassword = "" // パスワードは返さない

	return &m, nil
}

func (repository *userRepository) Update(ctx context.Context, m *model.User) (*model.User, error) {
	err := repository.Conn.Model(m).Update(m).Error
	return m, err
}

func (repository *userRepository) Delete(ctx context.Context, id int) error {
	m := model.User{ID: id}
	// gormはDeletedAtが含まれている場合は論理削除を行う
	return repository.Conn.Delete(&m).Error
}
