package routes

import (
	"aviones/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HOLA DESDE AVIONES ✈️")
	})

	app.Get("/aviones", controllers.ObtenerAviones)

	app.Get("/aviones/:id", controllers.ObtenerAvionPorID)

	app.Get("/aviones/check/matricula/:matricula", controllers.CheckMatricula)

	app.Get("/modelos", controllers.ObtenerModelos)

	app.Post("/modelos/create", controllers.CrearNuevoModelo)

	app.Get("/modelos/:id", controllers.ObtenerModeloPorID)

	app.Get("/modelos/find/:tipo", controllers.ObtenerModeloPorTipo)
}
