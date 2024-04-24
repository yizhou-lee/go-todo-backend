package main

import (
	"log"
	"net/http"
	"todo-backend/configs"
	"todo-backend/models"
	"todo-backend/routes"
	"todo-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Read configuration
	viper.SetConfigFile("configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database and migrate the model
	db := configs.ConnectMySQL()
	configs.AutoMigrate(db, &models.Todo{})

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
