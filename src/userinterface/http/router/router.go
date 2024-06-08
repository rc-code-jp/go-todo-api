// Package router ルーティング定義
package router

import (
	"os"

	"go-todo-api/src/userinterface/http/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetRoutes Router設定。
func SetRoutes(e *echo.Echo, handler handler.AppHandler) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Static("/", "assets")

	// アクセス制限なし
	unauthenticatedGroup := e.Group("/api/v1")
	unauthenticatedGroup.GET("/hc", handler.HealthCheck)
	unauthenticatedGroup.POST("/users/images", handler.UploadImageFile)
	unauthenticatedGroup.POST("/users", handler.CreateUser)
	unauthenticatedGroup.POST("/login", handler.Login)

	// アクセス制限あり
	authenticatedGroup := e.Group("/api/v1")
	authenticatedGroup.Use(middleware.JWT([]byte(os.Getenv("JWT_SIGNING_KEY"))))
	authenticatedGroup.GET("/users/:id", handler.GetUser)
	authenticatedGroup.PUT("/users/:id", handler.UpdateUser)
	authenticatedGroup.DELETE("/users/:id", handler.DeleteUser)
}
