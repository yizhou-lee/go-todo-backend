package routes

import (
	"github.com/gin-gonic/gin"
	"todo-backend/controllers"
)

var todoController = &controllers.TodoController{}

// TodoRoute represents todo-related routes
func TodoRoute(router *gin.Engine) {
	router.GET("/todos", todoController.GetTodos)
	router.GET("/todo/:id", todoController.GetTodoByID)
	router.POST("/todo", todoController.CreateTodo)
	router.PUT("/todo/:id", todoController.UpdateTodo)
	router.DELETE("/todo/:id", todoController.DeleteTodoByID)
}
