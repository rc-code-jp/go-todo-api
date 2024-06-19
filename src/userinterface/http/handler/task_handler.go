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
	TaskHandler interface {
		GetAllTask(c echo.Context) error
		GetTask(c echo.Context) error
		CreateTask(c echo.Context) error
		UpdateTask(c echo.Context) error
		DeleteTask(c echo.Context) error
	}

	taskHandler struct {
		TaskUseCase usecase.TaskUseCase
	}
)

func NewTaskHandler(usecase usecase.TaskUseCase) TaskHandler {
	return &taskHandler{usecase}
}

func (handler *taskHandler) GetAllTask(c echo.Context) error {
	userId := c.Get("userId").(int)

	taskGroupId, err := strconv.Atoi(c.Param("taskGroupId"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "IDは数値で入力してください。")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	items, err := handler.TaskUseCase.GetAllTask(ctx, taskGroupId, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": items,
	})
}

func (handler *taskHandler) GetTask(c echo.Context) error {
	userId := c.Get("userId").(int)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "IDは数値で入力してください。")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	item, err := handler.TaskUseCase.GetTask(ctx, id, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"item": item,
	})
}

func (handler *taskHandler) CreateTask(c echo.Context) error {
	userId := c.Get("userId").(int)
	request := &request.CreateTaskRequest{}
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

	item, err := handler.TaskUseCase.CreateTask(
		ctx,
		&model.Task{
			Title: request.Title,
			UserID: userId,
			Date: 				 request.Date,
			Time:  			request.Time,
		},
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"item":  item,
	})
}

func (handler *taskHandler) UpdateTask(c echo.Context) error {
	userId := c.Get("userId").(int)

	request := &request.UpdateTaskRequest{}
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

	_, err := handler.TaskUseCase.UpdateTask(
		ctx,
		&model.Task{
			ID:            userId,
			Title:          request.Title,
			Date: 				request.Date,
			Time:  			request.Time,
		},
		userId,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (handler *taskHandler) DeleteTask(c echo.Context) error {
	userId := c.Get("userId").(int)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "IDは数値で入力してください。")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := handler.TaskUseCase.DeleteTask(ctx, id, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
