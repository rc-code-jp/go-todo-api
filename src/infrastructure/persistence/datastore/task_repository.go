package datastore

import (
	"context"
	"go-todo-api/src/domain/model"
	"go-todo-api/src/domain/repository"

	"github.com/jinzhu/gorm"
)

type taskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(Conn *gorm.DB) repository.TaskRepository {
	return &taskRepository{Conn}
}

func (repository *taskRepository) Create(ctx context.Context, m *model.Task) (*model.Task, error) {
	err := repository.Conn.Create(m).Error
	return m, err
}

func (repository *taskRepository) AllByTaskGroupId(ctx context.Context, taskGroupId int) (*[]model.Task, error) {
	var models []model.Task
	err := repository.Conn.Where("task_group_id = ?", taskGroupId).Find(&models).Error
	if err != nil {
		return nil, err
	}
	return &models, nil
}

func (repository *taskRepository) FindByID(ctx context.Context, id int) (*model.Task, error) {
	m := model.Task{ID: id}
	err := repository.Conn.Preload("Tasks").First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (repository *taskRepository) Update(ctx context.Context, m *model.Task) (*model.Task, error) {
	err := repository.Conn.Model(m).Update(m).Error
	return m, err
}

func (repository *taskRepository) DoneByTaskGroupId(ctx context.Context, taskGroupId int) (error) {
	// TODO: 書く
	return nil
}

func (repository *taskRepository) Delete(ctx context.Context, id int) error {
	m := model.Task{ID: id}
	// gormはDeletedAtが含まれている場合は論理削除を行う
	return repository.Conn.Delete(&m).Error
}
