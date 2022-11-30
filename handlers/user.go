package handlers

import (
	"github.com/daniilmikhaylov2005/crudTodo/models"
	"github.com/daniilmikhaylov2005/crudTodo/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type responseUser struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type responseToken struct {
	Token string `json:"token"`
}

func SignUp(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Can't create user with this data",
		})
	}

	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Username or password can't be empty",
		})
	}

	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: "Error while hashing password",
		})
	}

	user.Password = hashPassword
	user.Role = "user"

	id, err := repository.InsertUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: "Error while inserting user",
		})
	}

	return c.JSON(http.StatusOK, responseUser{
		ID:     id,
		Status: "User created",
	})
}

func SignIn(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Can't login with this data",
		})
	}

	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Username or password can't be empty",
		})
	}

	userDb, err := repository.FindUserByUsername(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: "Error while selecting user",
		})
	}

	if CheckPasswordHash(user.Password, userDb.Password) {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Wrong password",
		})
	}

	token, err := CreateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, responseToken{
		Token: token,
	})
}
