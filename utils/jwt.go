package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "your_secret_key" // Replace with your actual secret key

func GenerateToken(email string, userid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userid": userid,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
