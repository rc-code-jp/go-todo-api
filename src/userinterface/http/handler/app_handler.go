// Package handler UI層
package handler

// AppHandler 全てのHandlerのinterfaceを満たす。
type AppHandler interface {
	UserHandler
	HcHandler
}

// appHandler 構造体
type appHandler struct {
	UserHandler
	HcHandler
}

// NewAppHandler AppHandlerを生成
func NewAppHandler(
	userHandler UserHandler,
	hcHandler HcHandler,
) AppHandler {
	return &appHandler{userHandler, hcHandler}
}
