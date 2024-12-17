package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"reservas/routes"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	routes.Setup(app)

	app.Listen(":3005")
}