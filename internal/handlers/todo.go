package handlers

import (
	"net/http"

	"RekaKR/Rest_API_example/internal/db"
	"RekaKR/Rest_API_example/internal/db/model"

	"github.com/gin-gonic/gin"
)

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, db.Todos)
}

func AddTodo(context *gin.Context) {
	var newTodo model.Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	db.Todos = append(db.Todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := db.GetTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func ToggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := db.GetTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}
