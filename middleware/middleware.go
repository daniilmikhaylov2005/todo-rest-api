package middleware

import (
	"errors"
	"fmt"
	"github.com/daniilmikhaylov2005/crudTodo/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func GetClaimsFromJWT(c echo.Context) (*models.UserClaims, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return &models.UserClaims{}, errors.New("Error while get jwt token from context variable")
	}

	claims, ok := user.Claims.(*models.UserClaims)
	if !ok {
		return &models.UserClaims{}, errors.New("Error while get claims from jwt token")
	}

	return claims, nil
}

func ParseToken(auth string, c echo.Context) (interface{}, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("%v\n", err)
	}

	jwtKey := os.Getenv("SIGNING_KEY")

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(jwtKey), nil
	}
	token, err := jwt.ParseWithClaims(auth, &models.UserClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid Token")
	}

	return token, nil
}
