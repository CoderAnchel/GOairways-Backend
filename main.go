package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"users/database"
	"users/routes"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	database.StartDB()
	defer database.CloseDB()
	routes.Setup(app)

	app.Listen(":3001")
}
