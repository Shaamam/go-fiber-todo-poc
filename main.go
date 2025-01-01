package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	v1 "go-fiber-todo-poc/api/v1"
	"go-fiber-todo-poc/config"
	_ "go-fiber-todo-poc/docs"
	"go-fiber-todo-poc/pkg/database"
)

// @title TODO API
// @version 1.0
// @description Provides API for DealerHUB
// @termsOfService http://swagger.io/terms/

// @contact.name Shaama

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.LoadConfig()
	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, OPTIONS, PUT, DELETE",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	v1.SetupRoutes(app)

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":8080")
}
