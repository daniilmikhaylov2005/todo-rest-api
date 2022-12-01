package main

import (
	"github.com/daniilmikhaylov2005/crudTodo/handlers"
	m "github.com/daniilmikhaylov2005/crudTodo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	api := e.Group("/api")

	auth := e.Group("/auth")
	auth.POST("/signin", handlers.SignIn)
	auth.POST("/signup", handlers.SignUp)

	config := middleware.JWTConfig{
		ParseTokenFunc: m.ParseToken,
	}

	api.Use(middleware.JWTWithConfig(config))

	api.GET("/todo", handlers.GetAllTodos)
	api.POST("/todo", handlers.CreateTodo)
	api.GET("/todo/:id", handlers.GetTodoById)
	api.PUT("/todo/:id", handlers.UpdateTodo)
	api.DELETE("/todo/:id", handlers.DeleteTodo)

	e.Logger.Fatal(e.Start(":8000"))
}
