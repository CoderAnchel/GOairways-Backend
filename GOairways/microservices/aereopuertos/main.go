package main

import (
	"aereopuertos/database"
	"aereopuertos/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	database.StartDB()
	defer database.CloseDB()

	app.Use(logger.New())
	routes.SetUp(app)

	app.Listen(":3003")
}
