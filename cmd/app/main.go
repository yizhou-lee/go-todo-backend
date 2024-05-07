package main

import (
	"log"
	"net/http"
	"todo-backend/configs"
	"todo-backend/internal/models"
	"todo-backend/internal/routes"
	"todo-backend/pkg/mysql"
	"todo-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Read configuration
	path := "./configs"
	cfg, err := configs.LoadConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database and migrate the model
	db := mysql.ConnectMySQL(cfg)
	mysql.AutoMigrate(db, &models.Todo{})

	router := gin.Default()

	// Routes
	routes.TodoRoute(router)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Page not found", nil))
	})

	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
