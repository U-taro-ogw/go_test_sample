package modules

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtTokenInterface interface {
	getEmail() string
}

type JwtToken struct {
	Email string
}

func (j JwtToken) getEmail() string {
	return j.Email
}


func GetJwtToken(j JwtTokenInterface) string {
	// ユーザ固有の情報もjwt tokenに含めた方が良さそう
	// 電子署名を取得する処理等...をメソッド分けする(した方が良いのか...?)

	// TODO 署名部分可変にするのは良くない
	jwtSignature := []byte(j.getEmail())
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(jwtSignature)
	return tokenString
}
