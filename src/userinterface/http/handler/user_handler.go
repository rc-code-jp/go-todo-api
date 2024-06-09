package handler

import (
	"context"
	"net/http"

	"go-todo-api/src/domain/model"
	"go-todo-api/src/usecase"
	"go-todo-api/src/userinterface/http/request"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler interface {
		CreateUser(c echo.Context) error
		Login(c echo.Context) error
		GetMe(c echo.Context) error
		UpdateUser(c echo.Context) error
		DeleteUser(c echo.Context) error
	}

	userHandler struct {
		UserUseCase usecase.UserUseCase
	}
)

func NewUserHandler(usecase usecase.UserUseCase) UserHandler {
	return &userHandler{usecase}
}

func (handler *userHandler) CreateUser(c echo.Context) error {
	request := &request.CreateUserRequest{}
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

	userID, token, err := handler.UserUseCase.CreateUser(
		ctx,
		request.Name,
		request.Email,
		request.Password,
		request.ImageFilePath,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":    userID,
		"token": token,
	})
}

func (handler *userHandler) Login(c echo.Context) error {
	request := &request.LoginRequest{}
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

	userID, token, err := handler.UserUseCase.Login(ctx, request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":    userID,
		"token": token,
	})
}

func (handler *userHandler) GetMe(c echo.Context) error {
	userId := c.Get("userId").(int)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := handler.UserUseCase.GetUser(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (handler *userHandler) UpdateUser(c echo.Context) error {
	userId := c.Get("userId").(int)

	request := &request.UpdateUserRequest{}
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

	_, err := handler.UserUseCase.UpdateUser(
		ctx,
		&model.User{
			ID:            userId,
			Name:          request.Name,
			ImageFilePath: request.ImageFilePath,
		},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (handler *userHandler) DeleteUser(c echo.Context) error {
	userId := c.Get("userId").(int)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := handler.UserUseCase.DeleteUser(ctx, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
