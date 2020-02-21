package handlers

import (
	"fmt"
	"github.com/U-taro-ogw/go_test_sample/auth_api/models"
	jwt "github.com/dgrijalva/jwt-go"
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
		fmt.Println("401エラーーーーー")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	} else {
		token := jwt.New(jwt.SigningMethodHS256)
		jwtToken, _ := token.SignedString([]byte("hoge"))
		c.JSON(http.StatusOK, gin.H{"jwt": jwtToken})
		return
	}
}
