package routes

import (
	"todo-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

var tc = &controllers.TodoController{}

// TodoRoute represents todo-related routes
func TodoRoute(r *gin.RouterGroup) {
	tr := r.Group("/todos")

	tr.GET("", tc.GetTodos)
	tr.POST("", tc.CreateTodo)
	tr.GET("/:id", tc.GetTodoByID)
	tr.PUT("/:id", tc.UpdateTodo)
	tr.DELETE("/:id", tc.DeleteTodoByID)
}
