package routes

import (
	"todo-backend/controllers"

	"github.com/gin-gonic/gin"
)

var todoController = &controllers.TodoController{}

// TodoRoute represents todo-related routes
func TodoRoute(router *gin.Engine) {
	router.GET("/todos", todoController.GetTodos)
	router.POST("/todos", todoController.CreateTodo)
	router.GET("/todos/:id", todoController.GetTodoByID)
	router.PUT("/todos/:id", todoController.UpdateTodo)
	router.DELETE("/todos/:id", todoController.DeleteTodoByID)
}
