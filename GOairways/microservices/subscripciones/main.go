package main

import (
	"github.com/gofiber/fiber/v2"
	"subscripciones/routes"
)

func main() {
	app := fiber.New()

	routes.Setup(app)
	app.Listen(":3009")
}
