package db

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB接続
func GetMysqlConnection() *gorm.DB {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	// 接続
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName,
	)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}

	// デバッグの設定
	debugMode, err := strconv.ParseBool(os.Getenv("DB_DEBUG_MODE"))
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(debugMode)

	return db
}

// DB接続
func NewDBConnection() *gorm.DB {
	return GetMysqlConnection()
}
