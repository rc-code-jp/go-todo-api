package handler

import (
	"context"
	"net/http"
	"strconv"

	"go-todo-api/src/domain/model"
	"go-todo-api/src/usecase"
	"go-todo-api/src/userinterface/http/request"

	"github.com/labstack/echo/v4"
)

type (
	TaskGroupHandler interface {
		GetAllTaskGroup(c echo.Context) error
		CreateTaskGroup(c echo.Context) error
		UpdateTaskGroup(c echo.Context) error
		DeleteTaskGroup(c echo.Context) error
	}

	taskGroupHandler struct {
		TaskGroupUseCase usecase.TaskGroupUseCase
	}
)

func NewTaskGroupHandler(usecase usecase.TaskGroupUseCase) TaskGroupHandler {
	return &taskGroupHandler{usecase}
}

func (handler *taskGroupHandler) GetAllTaskGroup(c echo.Context) error {
	userId := c.Get("userId").(int)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	items, err := handler.TaskGroupUseCase.GetAllTaskGroup(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": items,
	})
}

func (handler *taskGroupHandler) CreateTaskGroup(c echo.Context) error {
	userId := c.Get("userId").(int)
	request := &request.CreateTaskGroupRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	item, err := handler.TaskGroupUseCase.CreateTaskGroup(
		ctx,
		&model.TaskGroup{
			Name: request.Name,
			UserID: userId,
		},
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"item":  item,
	})
}

func (handler *taskGroupHandler) UpdateTaskGroup(c echo.Context) error {
	userId := c.Get("userId").(int)

	request := &request.UpdateTaskGroupRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err := handler.TaskGroupUseCase.UpdateTaskGroup(
		ctx,
		&model.TaskGroup{
			ID:            userId,
			Name:          request.Name,
		},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (handler *taskGroupHandler) DeleteTaskGroup(c echo.Context) error {
	userId := c.Get("userId").(int)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "IDは数値で入力してください。")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := handler.TaskGroupUseCase.DeleteTaskGroup(ctx, id, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
