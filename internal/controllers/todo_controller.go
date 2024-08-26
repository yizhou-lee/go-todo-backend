package controllers

import (
	"fmt"
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
//	@Success		200	{object}	utils.Response
//	@Router			/todos [get]
func (tc *TodoController) GetTodos(ctx *gin.Context) {
	todos, err := models.GetTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(fmt.Sprintf("Internal server error: %s", err), nil))
		return
	}
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
//	@Success		200	{object}	utils.Response
//	@Router			/todos/{id} [get]
func (tc *TodoController) GetTodoByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	todo, err := models.GetTodoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(fmt.Sprintf("Internal server error: %s", err), nil))
		return
	}
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
//	@Success		201		{object}	utils.Response
//	@Router			/todos [post]
func (tc *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(fmt.Sprintf("Internal server error: %s", err), nil))
		return
	}
	if err := models.CreateTodo(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(fmt.Sprintf("Internal server error: %s", err), nil))
		return
	}
	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Todo created successfully", nil))
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
//	@Success		200		{object}	utils.Response
//	@Router			/todos/{id} [put]
func (tc *TodoController) UpdateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(fmt.Sprintf("Internal server error: %s", err), nil))
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := models.UpdateTodoByID(id, &todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(fmt.Sprintf("Internal server error: %s", err), nil))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Todo updated successfully", nil))
}

// DeleteTodoByID DeleteTodo godoc
//
//	@Summary		Delete a todo
//	@Description	Delete an existing todo by its ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	utils.Response
//	@Router			/todos/{id} [delete]
func (tc *TodoController) DeleteTodoByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := models.DeleteTodoByID(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(fmt.Sprintf("Internal server error: %s", err), nil))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Todo deleted successfully", nil))
}
