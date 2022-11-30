package main

import (
	"github.com/daniilmikhaylov2005/crudTodo/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	e := echo.New()

	api := e.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/signin", handlers.SignIn)
	auth.POST("/signup", handlers.SignUp)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("%v\n", err)
	}

	jwtKey := os.Getenv("SIGNING_KEY")

	api.Use(middleware.JWT([]byte(jwtKey)))

	api.GET("/todo", handlers.GetAllTodos)
	api.POST("/todo", handlers.CreateTodo)
	api.GET("/todo/:id", handlers.GetTodoById)
	api.PUT("/todo/:id", handlers.UpdateTodo)
	api.DELETE("/todo/:id", handlers.DeleteTodo)

	e.Logger.Fatal(e.Start(":8000"))
}
