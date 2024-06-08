package main

import (
	"fmt"
	"os"

	"go-todo-api/src/infrastructure/db"
	"go-todo-api/src/interactor"
	"go-todo-api/src/userinterface/http/router"
	"go-todo-api/src/userinterface/http/validator"

	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

func main() {
	// Echoのインスタンス作成
	e := echo.New()

	// .envファイルの読み込み
	err := godotenv.Load()
  if err != nil {
    e.Logger.Fatal("Error loading .env file")
  }

	// DBコネクション作成
	conn := db.NewDBConnection()

	// DI
	interactor := interactor.NewInteractor(conn)

	// ルーティング
	e.Validator = validator.NewValidator()
	handler := interactor.NewAppHandler()
	router.SetRoutes(e, handler)

	if err := e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))); err != nil {
		e.Logger.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}
}
