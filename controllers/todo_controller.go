package controllers

import (
	"net/http"
	"strconv"
	"todo-backend/models"
	"todo-backend/utils"

	"github.com/gin-gonic/gin"
)

// TodoController represents a todo-related controller
type TodoController struct{}

func (tc *TodoController) GetTodos(c *gin.Context) {
	todos := models.GetTodos()
	c.JSON(http.StatusOK, utils.SuccessResponse("Todos fetched successfully", todos))
}

func (tc *TodoController) GetTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := models.GetTodoByID(id)
	c.JSON(http.StatusOK, utils.SuccessResponse("Todo fetched successfully", todo))
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		return
	}
	models.CreateTodo(&todo)
	c.JSON(http.StatusCreated, utils.SuccessResponse("Todo created successfully", todo))
}

func (tc *TodoController) UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		return
	}
	models.UpdateTodoByID(id, &todo)
	c.JSON(http.StatusOK, utils.SuccessResponse("Todo updated successfully", todo))
}

func (tc *TodoController) DeleteTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := models.DeleteTodoByID(id)
	c.JSON(http.StatusOK, utils.SuccessResponse("Todo deleted successfully", todo))
}
