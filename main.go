package main

import (
	"log"
	"net/http"
	"todo-backend/configs"
	"todo-backend/models"
	"todo-backend/routes"
	"todo-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database and migrate
	db := configs.ConnectMySQL()
	configs.AutoMigrate(db, &models.Todo{})

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
