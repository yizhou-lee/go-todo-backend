package configs

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// ConnectMySQL connects to MySQL and returns a pointer to the database.
func ConnectMySQL() *gorm.DB {
	once.Do(func() {
		host := viper.GetString("mysql.host")
		port := viper.GetString("mysql.port")
		username := viper.GetString("mysql.username")
		password := viper.GetString("mysql.password")
		database := viper.GetString("mysql.database")

		fmt.Printf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local\n", username, password, host, port, database)
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
func AutoMigrate(db *gorm.DB, model interface{}) {
	err := db.AutoMigrate(&model)
	if err != nil {
		log.Fatal(err)
	}
}
