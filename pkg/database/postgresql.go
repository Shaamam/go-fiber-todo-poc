package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	config "go-fiber-todo-poc/config"
	"log"
	"net/url"
)

var Pool *pgxpool.Pool

func Init() {
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

	// Setting up the connection pool
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v\n", err)
	}

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	// Test the connection
	var greeting string
	err = Pool.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Println(greeting) // Should print "Hello, world!"
}

// Remember to close the pool when your application exits
func Close() {
	if Pool != nil {
		Pool.Close()
	}
}
