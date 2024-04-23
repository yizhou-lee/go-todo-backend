package configs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectMySQL connects to MySQL and returns a pointer to the database.
func ConnectMySQL() *gorm.DB {
	dsn := EnvMySQLURI()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQLDB!")
	return db
}

// DB is a pointer to the database.
var DB = ConnectMySQL()
