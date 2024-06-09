package datastore

import (
	"context"
	"go-todo-api/src/domain/model"
	"go-todo-api/src/domain/repository"

	"github.com/jinzhu/gorm"
)

type taskGroupRepository struct {
	Conn *gorm.DB
}

func NewTaskGroupRepository(Conn *gorm.DB) repository.TaskGroupRepository {
	return &taskGroupRepository{Conn}
}

func (repository *taskGroupRepository) Create(ctx context.Context, m *model.TaskGroup) (*model.TaskGroup, error) {
	err := repository.Conn.Create(m).Error
	return m, err
}

func (repository *taskGroupRepository) All(ctx context.Context, userId int) (*[]model.TaskGroup, error) {
	var models []model.TaskGroup
	err := repository.Conn.Where("user_id = ?", userId).Find(&models).Error
	if err != nil {
		return nil, err
	}
	return &models, nil
}

func (repository *taskGroupRepository) FindByID(ctx context.Context, id int) (*model.TaskGroup, error) {
	m := model.TaskGroup{ID: id}
	err := repository.Conn.Preload("Tasks").First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (repository *taskGroupRepository) Update(ctx context.Context, m *model.TaskGroup) (*model.TaskGroup, error) {
	err := repository.Conn.Model(m).Update(m).Error
	return m, err
}

func (repository *taskGroupRepository) Delete(ctx context.Context, id int) error {
	m := model.TaskGroup{ID: id}
	// gormはDeletedAtが含まれている場合は論理削除を行う
	return repository.Conn.Delete(&m).Error
}
