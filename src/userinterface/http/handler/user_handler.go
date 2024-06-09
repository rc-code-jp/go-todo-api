package handler

import (
	"context"
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"

	"go-todo-api/src/domain/model"
	"go-todo-api/src/usecase"
	"go-todo-api/src/userinterface/http/request"

	"github.com/labstack/echo/v4"
	"github.com/olahol/go-imageupload"
)

type (
	UserHandler interface {
		UploadImageFile(c echo.Context) error
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

// 画像ファイルアップロード
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

// ユーザー作成
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

// ログイン
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

// ユーザー取得
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

// ユーザー更新
func (handler *userHandler) UpdateUser(c echo.Context) error {
	userId := c.Get("userId").(int)

	request := &request.UpdateUserRequest{ID: userId}
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

// ユーザー削除
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
