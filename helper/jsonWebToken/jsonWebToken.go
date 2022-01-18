package jsonWebToken

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(claims jwt.MapClaims, tokenExpiredInSecond int) (*string, error) {
	var err error

	claims["exp"] = time.Now().Add(time.Second * time.Duration(tokenExpiredInSecond)).Unix()
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtClaims.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func ParseJWT(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
