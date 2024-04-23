package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-backend/models"
)

// TodoController represents a todo-related controller
type TodoController struct{}

func (tc *TodoController) GetTodos(c *gin.Context) {
	todos := models.GetTodos()
	c.JSON(http.StatusOK, todos)
}

func (tc *TodoController) GetTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := models.GetTodoByID(id)
	c.JSON(http.StatusOK, todo)
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		return
	}
	models.CreateTodo(&todo)
	c.JSON(http.StatusCreated, todo)
}

func (tc *TodoController) UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		return
	}
	models.UpdateTodoByID(id, &todo)
	c.JSON(http.StatusOK, todo)
}

func (tc *TodoController) DeleteTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := models.DeleteTodoByID(id)
	c.JSON(http.StatusOK, todo)
}
