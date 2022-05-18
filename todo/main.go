package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo represents data about a todo.
type todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

var todos = []todo{
	{ID: "1", Title: "Blue Train", Description: "John Coltrane", Status: true},
	{ID: "2", Title: "Jeru", Description: "Gerry Mulligan", Status: true},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Description: "Sarah Vaughan", Status: false},
}

// gettodos responds with the list of all todos as JSON.
func getTodos(c *gin.Context) {
	var output = make(map[string][]todo)
	output["todos"] = todos
	c.IndentedJSON(http.StatusOK, output)
}

// postTodos adds an todo from JSON received in the request body.
func postTodos(c *gin.Context) {
	var newTodo todo

	// Call BindJSON to bind the received JSON to
	// newTodo.
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	// Add the new todo to the slice.
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", postTodos)

	router.Run("localhost:8080")
}
