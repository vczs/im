package help

import (
	"im/define"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Uid  string
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(uid, name string, second int) (string, error) {
	uc := UserClaim{
		Uid:              uid,
		Name:             name,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(second)))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
