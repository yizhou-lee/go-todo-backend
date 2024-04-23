package main

import (
	"log"
	"net/http"
	"todo-backend/configs"
	"todo-backend/routes"
	"todo-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to MySQLDB
	configs.ConnectMySQL()

	// Routes
	routes.TodoRoute(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Route not found", nil))
	})

	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
