package routes

import (
	"todo-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

var tc = &controllers.TodoController{}

// TodoRoute represents todo-related routes
func TodoRoute(r *gin.RouterGroup) {
	r.GET("/todos", tc.GetTodos)
	r.POST("/todos", tc.CreateTodo)
	r.GET("/todos/:id", tc.GetTodoByID)
	r.PUT("/todos/:id", tc.UpdateTodo)
	r.DELETE("/todos/:id", tc.DeleteTodoByID)
}
