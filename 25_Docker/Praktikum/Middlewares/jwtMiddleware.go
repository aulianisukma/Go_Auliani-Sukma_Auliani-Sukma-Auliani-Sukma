package middleware

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func CreateJWTToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"name":   name,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}