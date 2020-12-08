package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-api/models"
	"github.com/krittawatcode/go-todo-api/repository"
)

// TODOController is main interface for collect all method
type TODOController interface {
	GetToDos(c *gin.Context)
	CreateATodo(c *gin.Context)
	GetATodo(c *gin.Context)
	UpdateATodo(c *gin.Context)
	DeleteATodo(c *gin.Context)
}

type todoController struct {
}

// TDC is a global controller
var TDC TODOController

// TODOHandler use for init controller
func TODOHandler() TODOController {
	return &todoController{}
}

// GetToDos use to GET all todos
func (t *todoController) GetToDos(c *gin.Context) {

	var todo []models.Todo
	err := repository.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// CreateATodo use for CREATE Todo
func (t *todoController) CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := repository.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Get a particular Todo with id
func (t *todoController) GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := repository.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Update an existing Todo
func (t *todoController) UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := repository.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = repository.UpdateATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Delete a Todo
func (t *todoController) DeleteATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := repository.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
