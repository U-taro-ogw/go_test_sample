package handlers

import (
	"fmt"
	"github.com/U-taro-ogw/go_test_sample/auth_api/models"
	"github.com/gin-gonic/gin"
	//"github.com/gomodule/redigo/redis"
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

	h.Db.NewRecord(newUser)
	fmt.Println(h.Db.Create(&newUser))


	c.JSON(http.StatusCreated, gin.H{"message": "signup"})
}

//func (h *UserHandler) Signin(c *gin.Context) {
//	var userParam models.User
//	var findUser models.User
//	c.BindJSON(&userParam)
//
//	if err := h.Db.Where("email = ? AND password = ?", userParam.Email, userParam.Password).First(&findUser).Error; gorm.IsRecordNotFoundError(err) {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//	} else {
//		jwtToken := modules.GetTokenHandler()
//		modules.SetRedis(h.Redis, jwtToken, "111")
//		c.JSON(http.StatusOK, gin.H{"jwt": jwtToken})
//	}
//}
