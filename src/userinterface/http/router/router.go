package router

import (
	"go-todo-api/src/userinterface/http/handler"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JWTのClaims
type jwtClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

// SetRoutes Router設定。
func SetRoutes(e *echo.Echo, handler handler.AppHandler) {
	secret := os.Getenv("JWT_SIGNING_KEY")

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
	unauthenticatedGroup.POST("/users", handler.CreateUser)
	unauthenticatedGroup.POST("/login", handler.Login)

	// アクセス制限あり
	authenticatedGroup := e.Group("/api/v1")
	authenticatedGroup.Use(echoJwt.WithConfig(
		echoJwt.Config{
			SigningKey: []byte(secret),
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(jwtClaims)
			},
		},
	))
	authenticatedGroup.Use(extractUserID)
	// ユーザー
	authenticatedGroup.GET("/me", handler.GetMe)
	authenticatedGroup.DELETE("/me", handler.DeleteUser)
	authenticatedGroup.PUT("/me", handler.UpdateUser)
	// タスクグループ
	authenticatedGroup.GET("/task-groups", handler.GetAllTaskGroup)
	authenticatedGroup.POST("/task-groups", handler.CreateTaskGroup)
	authenticatedGroup.PUT("/task-groups/:id", handler.UpdateTaskGroup)
	authenticatedGroup.DELETE("/task-groups/:id", handler.DeleteTaskGroup)
}

// IdをContextに入れる
func extractUserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*jwtClaims)
			userId := claims.ID
			c.Set("userId", userId)
			return next(c)
	}
}
