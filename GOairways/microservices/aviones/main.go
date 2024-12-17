package main

import (
	"aviones/database"
	"aviones/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Database
	database.StartDB()
	defer database.CloseDB()

	// Middleware
	app.Use(logger.New())

	// Routes
	routes.Setup(app)

	app.Listen(":3002")
}
