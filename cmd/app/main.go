package main

import (
	"log"
	"net/http"
	"todo-backend/configs"
	_ "todo-backend/docs"
	"todo-backend/internal/models"
	"todo-backend/internal/routes"
	"todo-backend/pkg/mysql"
	"todo-backend/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Todo Backend API
// @description	This is a sample server for a todo backend.
// @version		1.0
// @host			localhost:8080
// @BasePath		/
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

	r := gin.Default()

	// Routes
	routes.TodoRoute(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Page not found", nil))
	})
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
