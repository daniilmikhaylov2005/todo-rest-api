package main

import (
	"github.com/daniilmikhaylov2005/crudTodo/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	api := e.Group("/api")
	api.GET("/todo", handlers.GetAllTodos)
	api.POST("/todo", handlers.CreateTodo)
	api.GET("/todo/:id", handlers.GetTodoById)
	api.PUT("/todo/:id", handlers.UpdateTodo)
	api.DELETE("/todo/:id", handlers.DeleteTodo)
	e.Logger.Fatal(e.Start(":8000"))
}
