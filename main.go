package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{

	{ID: "1", Item: "Clean your room", Completed: false},
	{ID: "2", Item: "Clean your balls", Completed: false},
	{ID: "3", Item: "Go out and play", Completed: false},
	{ID: "4", Item: "Study Goland", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func addTodo(context *gin.Context) {
	var newTodo Todo

	if error := context.BindJSON(&newTodo); error != nil {
		return
	}
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*Todo, error) {

	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")

}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	} else {
		context.IndentedJSON(http.StatusOK, todo)
	}

}

func patchTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	} else {

		todo.Completed = !todo.Completed

		context.IndentedJSON(http.StatusOK, todo)
	}

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", patchTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:8000")
}
