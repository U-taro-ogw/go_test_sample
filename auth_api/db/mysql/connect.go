package mysql

import (
	"github.com/U-taro-ogw/go_test_sample/auth_api/models"
	"github.com/jinzhu/gorm"
	"os"
)

func MysqlConnect() *gorm.DB {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_ROOT_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic("connection failed auth_db")
	}
	db.AutoMigrate(&models.User{})

	return db
}