package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/pkg/models"
)

func (a App) CreateTodo(c *gin.Context) {

	newTodo, err := models.ValidateCreateTodoRequest(c)

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err)
		return
	}

	err = a.todoSvc.CreateTodo(newTodo)

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err)
		return
	}

	c.String(http.StatusCreated, "")
}
func (a App) GetAllTodos(c *gin.Context) {

	list, err := a.todoSvc.GetAllTodos()

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": list,
	})
}

func (a App) DeleteTodos(c *gin.Context) {
	ids, err := models.ValidateDeletionRequest(c)

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err)
		return
	}
	err = a.todoSvc.DeleteTodos(ids)

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err)
		return
	}

	c.String(http.StatusOK, "deletion success")
}
