// Package handler UI層
package handler

import (
	"context"
	"crypto/sha1"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go-todo-api/src/domain/model"
	"go-todo-api/src/usecase"
	"go-todo-api/src/userinterface/http/request"

	"github.com/labstack/echo/v4"
	"github.com/olahol/go-imageupload"
)

type (
	// UserHandler interface
	UserHandler interface {
		UploadImageFile(c echo.Context) error
		CreateUser(c echo.Context) error
		Login(c echo.Context) error
		GetUser(c echo.Context) error
		UpdateUser(c echo.Context) error
		DeleteUser(c echo.Context) error
	}

	// userHandler 構造体
	userHandler struct {
		UserUseCase usecase.UserUseCase
	}
)

// NewUserHandler UserHandlerを生成。
func NewUserHandler(usecase usecase.UserUseCase) UserHandler {
	return &userHandler{usecase}
}

// UploadImageFile　ユーザープロフィール画像をアップロードし、画像ファイルのパスを返す。
func (handler *userHandler) UploadImageFile(c echo.Context) error {
	img, err := imageupload.Process(c.Request(), "ImageFile")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	h := sha1.Sum(thumb.Data)
	fileName := fmt.Sprintf("%s_%x.png", time.Now().Format("20060102150405"), h[:])
	thumb.Save("assets/images/" + fileName)

	return c.String(http.StatusOK, "images/"+fileName)
}

// CreateUser 登録
func (handler *userHandler) CreateUser(c echo.Context) error {
	request := new(request.CreateUserRequest)
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

// Login ログイン。ユーザーID、JWTトークンを返す。
func (handler *userHandler) Login(c echo.Context) error {
	request := new(request.LoginRequest)
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

// GetUser 詳細取得
func (handler *userHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "IDは数値で入力してください。")
	}

	request := &request.GetUserRequest{ID: id}
	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := handler.UserUseCase.GetUser(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser 更新
func (handler *userHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "IDは数値で入力してください。")
	}

	request := &request.UpdateUserRequest{ID: id}
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

	_, err = handler.UserUseCase.UpdateUser(
		ctx,
		&model.User{
			ID:            id,
			Name:          request.Name,
			Email:         request.Email,
			ImageFilePath: request.ImageFilePath,
		},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// DeleteUser 削除
func (handler *userHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "IDは数値で入力してください。")
	}

	request := request.DeleteUserRequest{ID: id}
	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := handler.UserUseCase.DeleteUser(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
