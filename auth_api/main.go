package main

import (
	"fmt"
	authDb "github.com/U-taro-ogw/go_test_sample/auth_api/db/mysql"
	//authenticationDb "github.com/U-taro-ogw/go_test_sample/auth_api/db/redis"
	"github.com/U-taro-ogw/go_test_sample/auth_api/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func GetMainEngine(userHandler handlers.UserHandler) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.POST("/signup", userHandler.Signup)
		v1.POST("/signin", userHandler.Signin)
	}
	return r
}

func main() {
	dbCon := authDb.MysqlConnect()
	//defer dbCon.Close()
	dbCon.LogMode(true)

	userHandler := handlers.UserHandler{Db: dbCon}
	fmt.Println("Hello, World!")
	GetMainEngine(userHandler).Run(":8083")
}
