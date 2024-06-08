// Package handler UI層
package handler

// AppHandler 全てのHandlerのinterfaceを満たす。
type AppHandler interface {
	UserHandler
}

// appHandler 構造体
type appHandler struct {
	UserHandler
}

// NewAppHandler AppHandlerを生成
func NewAppHandler(userHandler UserHandler) AppHandler {
	return &appHandler{userHandler}
}
