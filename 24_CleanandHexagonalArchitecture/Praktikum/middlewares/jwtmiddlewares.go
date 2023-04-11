package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type JWTS interface {
	CreateJWTToken(id uint, email string) (string, error)
}

type jwtS struct {
	issuer    string
	secretKey string
}

func NewJWTS() JWTS {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
		log.Println(err)
	}
	return &jwtS{
		issuer:    "qwerty",
		secretKey: os.Getenv("JWT_KEY"),
	}
}

func (j *jwtS) CreateJWTToken(id uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(10 * time.Minute).Unix(),
		"iss":   j.issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}