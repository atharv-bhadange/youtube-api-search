package cmd

import (
	"github.com/atharv-bhadange/youtube-api-search/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// InitApp initializes the fiber app
func InitApp() *fiber.App {
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET",
		AllowHeaders: "Origin, Content-Type, Accept, Accept-Language, Content-Length, Authorization, X-Api-Key",
	}))

	// Initialize routes
	routes.InitRoutes(app)

	return app
}
