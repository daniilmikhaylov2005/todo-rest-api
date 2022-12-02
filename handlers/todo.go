package handlers

import (
	"fmt"
	"github.com/daniilmikhaylov2005/crudTodo/middleware"
	"github.com/daniilmikhaylov2005/crudTodo/models"
	"github.com/daniilmikhaylov2005/crudTodo/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

type response struct {
	Status string `json:"status"`
}

func GetAllTodos(c echo.Context) error {
	var todos []models.Todo

	todos = repository.GetAllTodos()

	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	var todo models.Todo

	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status: "Can't create todo with this data",
		})
		return err
	}

	// check role in jwt token claims
	claims, err := middleware.GetClaimsFromJWT(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: fmt.Sprintf("Error with jwt claims, %v", err),
		})
	}

	if claims.Role != "admin" {
		return c.JSON(http.StatusBadRequest, response{
			Status: fmt.Sprintf("Only admin can create todo!"),
		})
	}

	recievedId := repository.CreateTodo(todo)

	todo.ID = recievedId

	return c.JSON(http.StatusCreated, todo)
}

func GetTodoById(c echo.Context) error {
	var todo models.Todo

	textId := c.Param("id")
	id, err := strconv.Atoi(textId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Unable to convert id to string",
		})
	}

	todo, err = repository.GetTodoById(id)

	if err != nil {
		return errorResponse(c, err)
	}

	return c.JSON(http.StatusOK, todo)

}

func UpdateTodo(c echo.Context) error {
	var todo models.Todo

	textId := c.Param("id")
	id, err := strconv.Atoi(textId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Unable to convert id to string",
		})
	}

	// check role in jwt token claims
	claims, err := middleware.GetClaimsFromJWT(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: fmt.Sprintf("Error with jwt claims, %v", err),
		})
	}

	if claims.Role != "admin" {
		return c.JSON(http.StatusBadRequest, response{
			Status: fmt.Sprintf("Only admin can update todo!"),
		})
	}

	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status: "Can't update todo with this data",
		})
		return err
	}

	// get todo to check json data
	recievedTodo, err := repository.GetTodoById(id)

	if err != nil {
		return errorResponse(c, err)

	}

	// service
	if strings.TrimSpace(todo.Title) == "" {
		todo.Title = recievedTodo.Title
	}
	if todo.Done == false {
		todo.Done = recievedTodo.Done
	}

	err = repository.UpdateTodo(id, todo)

	if err != nil {
		return errorResponse(c, err)
	}

	todo.ID = id

	return c.JSON(http.StatusAccepted, todo)
}

func DeleteTodo(c echo.Context) error {
	textId := c.Param("id")
	id, err := strconv.Atoi(textId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Unable to convert id to string",
		})
	}

	// check role in jwt token claims
	claims, err := middleware.GetClaimsFromJWT(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: fmt.Sprintf("Error with jwt claims, %v", err),
		})
	}

	if claims.Role != "admin" {
		return c.JSON(http.StatusBadRequest, response{
			Status: fmt.Sprintf("Only admin can delete todo!"),
		})
	}

	deletedId, err := repository.DeleteTodo(id)

	if err != nil {
		return errorResponse(c, err)
	}

	msg := fmt.Sprintf("todo with id %d deleted", deletedId)
	return c.JSON(http.StatusOK, response{
		Status: msg,
	})
}
