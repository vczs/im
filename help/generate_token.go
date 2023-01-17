package help

import (
	"im/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Uid   string
	Email string
	jwt.RegisteredClaims
}

func GenerateToken(uid, email string, second int) (string, error) {
	uc := UserClaim{
		Uid:              uid,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(second)))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
