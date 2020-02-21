package modules

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtToken interface {
	GetJwtToken() string
}

func GetJwtToken() string {
	// ユーザ固有の情報もjwt tokenに含めた方が良さそう
	// 電子署名を取得する処理等...をメソッド分けする(した方が良いのか...?)
	jwtSignature := []byte("hoge")
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(jwtSignature)
	return tokenString
}
