package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rizkyitaji/foodmarket-backend-go/database"
	"github.com/rizkyitaji/foodmarket-backend-go/routes"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env variables")
	}

	// Connect DB
	database.Connect()

	// Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Run server
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
