package middleware

import (
	"errors"
	"fmt"
	"github.com/daniilmikhaylov2005/crudTodo/handlers"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

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
	token, err := jwt.ParseWithClaims(auth, &handlers.UserClaims{}, keyFunc)
	if err != nil {
		fmt.Println(1.5)
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid Token")
	}

	claims := token.Claims

	c.JSON(http.StatusOK, claims)
	return token, nil
}
