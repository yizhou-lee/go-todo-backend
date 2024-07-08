package controllers

import (
	"net/http"
	"strconv"
	"todo-backend/internal/models"
	"todo-backend/utils"

	"github.com/gin-gonic/gin"
)

// TodoController represents a todo-related controller
type TodoController struct{}

// GetTodos godoc
//
//	@Summary		Get list of todos
//	@Description	Get a list of all todos
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/todos [get]
func (tc *TodoController) GetTodos(ctx *gin.Context) {
	todos := models.GetTodos()
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Todos fetched successfully", todos))
}

// GetTodoByID godoc
//
//	@Summary		Get a todo by ID
//	@Description	Get a single todo by its ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	map[string]interface{}
//	@Router			/todos/{id} [get]
func (tc *TodoController) GetTodoByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	todo := models.GetTodoByID(id)
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Todo fetched successfully", todo))
}

// CreateTodo godoc
//
//	@Summary		Create a new todo
//	@Description	Create a new todo with the given details
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		models.Todo	true	"Create Todo"
//	@Success		201		{object}	map[string]interface{}
//	@Router			/todos [post]
func (tc *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	err := ctx.BindJSON(&todo)
	if err != nil {
		return
	}
	models.CreateTodo(&todo)
	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Todo created successfully", todo))
}

// UpdateTodo godoc
//
//	@Summary		Update an existing todo
//	@Description	Update an existing todo by its ID with new details
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"Todo ID"
//	@Param			todo	body		models.Todo	true	"Update Todo"
//	@Success		200		{object}	map[string]interface{}
//	@Router			/todos/{id} [put]
func (tc *TodoController) UpdateTodo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var todo models.Todo
	err := ctx.BindJSON(&todo)
	if err != nil {
		return
	}
	models.UpdateTodoByID(id, &todo)
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Todo updated successfully", todo))
}

// DeleteTodo godoc
//
//	@Summary		Delete a todo
//	@Description	Delete an existing todo by its ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Todo ID"
//	@Success		204	"Todo deleted successfully"
//	@Router			/todos/{id} [delete]
func (tc *TodoController) DeleteTodoByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	todo := models.DeleteTodoByID(id)
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Todo deleted successfully", todo))
}
