package modules

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// 関数は1つ以上の引数を取るべきだ
// って論があったはず
func GenerateJwtToken() string {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)
	// 電子署名
	tokenString, _ := token.SignedString([]byte("hoge"))

	return tokenString
}
