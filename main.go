package main

import (
	"log"
	"net/http"
	"todo-backend/configs"
	"todo-backend/models"
	"todo-backend/routes"
	"todo-backend/utils"

	"todo-backend/pkg/sql"

	"github.com/gin-gonic/gin"
)

func main() {
	// Read configuration
	configs.LoadConfig()

	// Connect to database and migrate the model
	db := sql.ConnectMySQL()
	sql.AutoMigrate(db, &models.Todo{})

	router := gin.Default()
	// Routes
	routes.TodoRoute(router)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Page not found", nil))
	})

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
