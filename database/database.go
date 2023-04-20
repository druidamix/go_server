// Database package manages the database
package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/druidamix/go_server/config"
	"github.com/druidamix/go_server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// Connect function
func Connect() {
	p := config.Config("DB_PORT")
	// because our config function returns a string, we are parsing our      str to int here
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("hello")
	log.Println("running migrations")
  err = db.AutoMigrate(&model.User{})

  if err != nil {
    fmt.Errorf("Error migrating")
  }
	DB = Dbinstance{
		Db: db,
	}
}
