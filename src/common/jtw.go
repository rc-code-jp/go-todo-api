package common

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTのClaims
type jwtClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

// JWTトークンを生成
func CreateJwtToken(userId int) (string, error) {
	// 鍵となる文字列
	secret := os.Getenv("JWT_SIGNING_KEY")

	claims := &jwtClaims{
		ID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))

	return t, err
}
