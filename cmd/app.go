package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// InitApp initializes the fiber app
func InitApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET",
		AllowHeaders: "Origin, Content-Type, Accept, Accept-Language, Content-Length, Authorization, X-Api-Key",
	}))

	return app
}
