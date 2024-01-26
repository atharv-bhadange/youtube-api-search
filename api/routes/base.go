package routes

import (
	C "github.com/atharv-bhadange/youtube-api-search/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	// Health check route
	app.Get("/", C.HealthCheck)

	// Search route
	app.Get("/search", C.GetVideos)
}
