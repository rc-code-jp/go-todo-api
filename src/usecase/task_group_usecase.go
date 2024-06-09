package usecase

import (
	"context"
	"errors"

	"go-todo-api/src/domain/model"
	"go-todo-api/src/domain/repository"
)

type TaskGroupUseCase interface {
	GetAllTaskGroup(ctx context.Context, taskGroupId int) (*[]model.TaskGroup, error)
	GetTaskGroup(ctx context.Context, id int, userId int) (*model.TaskGroup, error)
	CreateTaskGroup(ctx context.Context, m *model.TaskGroup) (*model.TaskGroup, error)
	UpdateTaskGroup(ctx context.Context, m *model.TaskGroup) (*model.TaskGroup, error)
	DeleteTaskGroup(ctx context.Context, id int, userId int) error
}

type taskGroupUseCase struct {
	repository.TaskGroupRepository
}

func NewTaskGroupUseCase(repository repository.TaskGroupRepository) TaskGroupUseCase {
	return &taskGroupUseCase{repository}
}

// GetAllTaskGroup get all task group
func (u *taskGroupUseCase) GetAllTaskGroup(ctx context.Context, userId int) (*[]model.TaskGroup, error) {
	return u.TaskGroupRepository.All(ctx, userId)
}

// GetTaskGroup get task group
func (u *taskGroupUseCase) GetTaskGroup(ctx context.Context, id int, userId int) (*model.TaskGroup, error) {
	model, err := u.TaskGroupRepository.FindByID(ctx, id)
	// ユーザーが一致するかどうか確認
	if err != nil || model.UserID != userId {
		return nil, errors.New("task group not found")
	}
	return model, err
}

// CreateTaskGroup create task group
func (u *taskGroupUseCase) CreateTaskGroup(ctx context.Context, m *model.TaskGroup) (*model.TaskGroup, error) {
	return u.TaskGroupRepository.Create(ctx, m)
}

// UpdateTaskGroup update task group
func (u *taskGroupUseCase) UpdateTaskGroup(ctx context.Context, m *model.TaskGroup) (*model.TaskGroup, error) {
	// ユーザーが一致するかどうか確認
	if m.UserID != m.UserID {
		return nil, errors.New("task group not found")
	}
	return u.TaskGroupRepository.Update(ctx, m)
}

// DeleteTaskGroup delete task group
func (u *taskGroupUseCase) DeleteTaskGroup(ctx context.Context, id int, userId int) error {
	// ユーザーが一致するかどうか確認
	model, err := u.TaskGroupRepository.FindByID(ctx, id)
	if err != nil || model.UserID != userId {
		return errors.New("task group not found")
	}

	return u.TaskGroupRepository.Delete(ctx, id)
}

