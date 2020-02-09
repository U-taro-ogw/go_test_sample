package modules

import (
	jwt "github.com/dgrijalva/jwt-go"
)

func GetTokenHandler() string {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)
	// 電子署名
	tokenString, _ := token.SignedString([]byte("hoge"))

	return tokenString
}
