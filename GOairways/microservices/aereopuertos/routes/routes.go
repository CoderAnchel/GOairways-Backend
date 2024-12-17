package routes

import (
	"aereopuertos/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Get("/", controllers.Hello)

	app.Get("/aereopuertos", controllers.ObtenerAereopuertos)
}
