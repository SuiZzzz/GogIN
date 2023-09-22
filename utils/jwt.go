package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 密钥
const key = "d4qw8416fa"

type Claims struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}

// GenerateToken 签发Token
func GenerateToken(id uint, username string) (string, error) {
	claim := Claims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "GoGin_Server",                                     // 签发者
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),    // 最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
		},
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := withClaims.SignedString([]byte(key))
	return token, err
}
