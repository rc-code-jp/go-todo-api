// Package handler UI層
package handler

import (
	"github.com/labstack/echo/v4"
)

type (
	// UserHandler interface
	HcHandler interface {
		HealthCheck(c echo.Context) error
	}

	// userHandler 構造体
	hcHandler struct {
	}
)

// NewUserHandler UserHandlerを生成。
func NewHcHandler() HcHandler {
	return &hcHandler{}
}

// ヘルスチェック
func (handler *hcHandler) HealthCheck(c echo.Context) error {
	return c.String(200, "OK")
}
