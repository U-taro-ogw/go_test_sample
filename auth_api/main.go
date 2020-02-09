package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/U-taro-ogw/go_test_sample/auth_api/db"
	//authenticationDb "github.com/U-taro-ogw/go_test_sample/auth_api/db/redis"
	"github.com/U-taro-ogw/go_test_sample/auth_api/handlers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("Hello, World!")

	dbCon := db.MysqlConnect()
	defer dbCon.Close()
	dbCon.LogMode(true)

	redisCon := db.RedisConnect()

	userHandler := handlers.UserHandler{
		Db: dbCon,
		Redis: redisCon,
	}
	d := gin.Default()
	d.POST("/signup", userHandler.Signup)
	//d.POST("/signin", userHandler.Signin)

	d.Run(":8083")
}
