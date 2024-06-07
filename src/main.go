package main

import (
	"go-todo-api/src/infrastructure/db"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// loggerの初期化
	// ToDo: ログのライブラリを初期化する（選定から）

	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		fmt.Println(".envファイルの読み込みに失敗しました")
		return
	}

	// DB接続
	db.InitializeDB()

	// echoを初期化
	// apiClient := presentation.NewApiClient()
	// apiClient.RegisterRoute()
	// apiClient.Start()
}
