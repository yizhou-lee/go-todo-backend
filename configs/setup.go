package configs

import (
	"log"
	"sync"

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
		dsn := EnvMySQLURI()
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connected to MySQLDB!")
	})
	return db
}

// DB is a pointer to the database.
var DB = ConnectMySQL()

// AutoMigrate migrates the model to the database.
func AutoMigrate(db *gorm.DB, model interface{}) {
	err := db.AutoMigrate(&model)
	if err != nil {
		log.Fatal(err)
	}
}
