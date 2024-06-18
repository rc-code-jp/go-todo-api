package usecase

import (
	"context"
	"errors"

	"go-todo-api/src/domain/model"
	"go-todo-api/src/domain/repository"
)

type TaskUseCase interface {
	GetAllTask(ctx context.Context, taskGroupId int, userId int) (*[]model.Task, error)
	GetTask(ctx context.Context, id int, userId int) (*model.Task, error)
	CreateTask(ctx context.Context, m *model.Task) (*model.Task, error)
	UpdateTask(ctx context.Context, m *model.Task, userId int) (*model.Task, error)
	DeleteTask(ctx context.Context, id int, userId int) error
}

type taskUseCase struct {
	repository.TaskRepository
	repository.TaskGroupRepository
}

func NewTaskUseCase(repository repository.TaskRepository, taskGroupRepository repository.TaskGroupRepository) TaskUseCase {
	return &taskUseCase{repository, taskGroupRepository}
}

func (u *taskUseCase) GetAllTask(ctx context.Context, taskGroupId int, userId int) (*[]model.Task, error) {
	model, err := u.TaskGroupRepository.FindByID(ctx, taskGroupId)
	// ユーザーが一致するかどうか確認
	if err != nil || model.UserID != userId {
		return nil, errors.New("task group not found")
	}
	return u.TaskRepository.AllByTaskGroupId(ctx, taskGroupId)
}

func (u *taskUseCase) GetTask(ctx context.Context, id int, userId int) (*model.Task, error) {
	model, err := u.TaskRepository.FindByID(ctx, id)
	// ユーザーが一致するかどうか確認
	if err != nil || model.UserID != userId {
		return nil, errors.New("task not found")
	}
	return model, err
}

func (u *taskUseCase) CreateTask(ctx context.Context, m *model.Task) (*model.Task, error) {
	return u.TaskRepository.Create(ctx, m)
}

func (u *taskUseCase) UpdateTask(ctx context.Context, m *model.Task, userId int) (*model.Task, error) {
	// ユーザーが一致するかどうか確認
	if m.UserID != userId {
		return nil, errors.New("task not found")
	}
	return u.TaskRepository.Update(ctx, m)
}

func (u *taskUseCase) DeleteTask(ctx context.Context, id int, userId int) error {
	// ユーザーが一致するかどうか確認
	model, err := u.TaskRepository.FindByID(ctx, id)
	if err != nil || model.UserID != userId {
		return errors.New("task group not found")
	}

	return u.TaskRepository.Delete(ctx, id)
}

