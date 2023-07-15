package main

import (
	"RekaKR/Rest_API_example/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/todos", handlers.GetTodos)
	router.POST("/todos", handlers.AddTodo)
	router.GET("/todos/:id", handlers.GetTodo)
	router.PATCH("/todos/:id", handlers.ToggleTodoStatus)

	router.Run("localhost:8080")
}
