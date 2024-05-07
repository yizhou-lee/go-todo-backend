package mysql

import (
	"fmt"
	"log"
	"sync"
	"todo-backend/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// ConnectMySQL connects to MySQL and returns a pointer to the database.
func ConnectMySQL(cfg *configs.Config) *gorm.DB {
	once.Do(func() {
		host := cfg.MySQL.Host
		port := cfg.MySQL.Port
		username := cfg.MySQL.Username
		password := cfg.MySQL.Password
		database := cfg.MySQL.Database

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connected to MySQLDB!")
	})
	return db
}

// GetDB returns the database connection.
func GetDB() *gorm.DB {
	return db
}

// AutoMigrate migrates the model to the database.
func AutoMigrate(db *gorm.DB, model ...interface{}) {
	err := db.AutoMigrate(model...)
	if err != nil {
		log.Printf("Failed to migrate model: %v", err)
	}
}
