package handlers

import (
	"fmt"
	"github.com/daniilmikhaylov2005/crudTodo/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func errorResponse(c echo.Context, err error) error {
	if err.Error() == "sql: no rows in result set" {
		return c.JSON(http.StatusNotFound, response{
			Status: "todo does not exist",
		})
	}
	status := fmt.Sprintf("%v", err)
	return c.JSON(http.StatusNotFound, response{
		Status: status,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(user models.User) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}
	claims := &models.UserClaims{
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringKey := os.Getenv("SIGNING_KEY")
	byteKey := []byte(stringKey)

	stringToken, err := token.SignedString(byteKey)
	return stringToken, err
}
