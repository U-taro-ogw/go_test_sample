package main

import (
	"fmt"
	//"github.com/U-taro-ogw/auth_api/src/db"
	//"github.com/U-taro-ogw/auth_api/src/handlers"
	"github.com/gin-gonic/gin"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("Hello, World!")

	//dbCon := db.DbConnect()
	//defer dbCon.Close()
	//dbCon.LogMode(true)
	//
	//redisCon := db.RedisConnect()
	//
	//userHandler := handlers.UserHandler{
	//	Db: dbCon,
	//	Redis: redisCon,
	//}
	//
	d := gin.Default()
	//d.POST("/signup", userHandler.Signup)
	//d.POST("/signin", userHandler.Signin)

	d.Run(":8083")
}
