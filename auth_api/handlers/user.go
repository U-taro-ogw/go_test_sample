package handlers

import (
	"fmt"
	"github.com/U-taro-ogw/go_test_sample/auth_api/db/redis"
	"github.com/U-taro-ogw/go_test_sample/auth_api/models"
	"github.com/U-taro-ogw/go_test_sample/auth_api/modules"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type UserHandler struct {
	Db *gorm.DB
}

func (h *UserHandler) Signup(c *gin.Context) {
	newUser := models.User{}
	err := c.BindJSON(&newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := validator.New()
	err = v.Struct(newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.Db.Create(&newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "signup"})
	return
}

func (h *UserHandler) Signin(c *gin.Context) {
	var userParam models.User
	var findUser models.User
	err := c.BindJSON(&userParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if err := h.Db.Where("email = ? AND password = ?", userParam.Email, userParam.Password).First(&findUser).Error; gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	} else {
		// jwt token生成
		j := modules.JwtToken{findUser.Email}
		jwtToken := modules.GetJwtToken(j)

		// redis保存
		r := redis.RedisConnect()
		modules.RedisSet(r, jwtToken, "111")

		fmt.Println("ログイン成功")
		c.JSON(http.StatusOK, gin.H{"jwt_token": jwtToken})
		return
	}
}
