package handler

import (
	"github.com/labstack/echo/v4"
)

type (
	HcHandler interface {
		HealthCheck(c echo.Context) error
	}

	hcHandler struct {
	}
)

func NewHcHandler() HcHandler {
	return &hcHandler{}
}

// ヘルスチェック
func (handler *hcHandler) HealthCheck(c echo.Context) error {
	return c.String(200, "OK")
}
