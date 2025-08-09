package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyitaji/foodmarket-backend-go/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/users", handlers.GetUsers)
	api.Post("/users", handlers.CreateUser)
}
