package database

import (
	"fmt"
	"go-fiber-todo-poc/config"
	"go-fiber-todo-poc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
)

// DB is the global database connection
var DB *gorm.DB

func ConnectDB() {
	dbUser := config.AppConfig.Database.User
	dbPassword := config.AppConfig.Database.Password
	dbName := config.AppConfig.Database.DBName
	dbHost := config.AppConfig.Database.Host
	dbPort := config.AppConfig.Database.Port

	// URL encode the username and password
	encodedUser := url.QueryEscape(dbUser)
	encodedPassword := url.QueryEscape(dbPassword)

	// Construct the connection string
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		encodedUser, encodedPassword, dbHost, dbPort, dbName,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Todo{})

	DB = db
}
