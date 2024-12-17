package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"rutas/database"
	"rutas/routes"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	routes.Setup(app)
	database.StartDB()
	defer database.CloseDB()
	app.Listen(":3006")
}
