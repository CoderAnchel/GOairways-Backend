package main

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	app.Get("GOairways/API/", hello)

	app.Get("GOairways/API/routes", getRoutes)

	app.Get("GOairways/API/routes/active", getActiveRoutes)

	app.Get("GOairways/API/routes/inactive", getInactiveRoutes)

	app.Get("GOairways/API/routes/accept/:id", activateRoute)

	app.Post("GOairways/API/routes/create", createNewRoute)

	app.Get("GOairways/API/aereopuertos", getAereopuertos)

	app.Get("GOairways/API/aviones/", getAviones)

	app.Get("GOairways/API/aviones/:id", GetAvionPorID)

	app.Get("GOairways/API/aviones/check/matricula/:matricula", checkMatricula)

	app.Get("/GOairways/API/modelos", getModelos)

	app.Get("GOairways/API/modelos/:id", getModeloById)

	app.Get("GOairways/API/modelos/find/:tipo", getModeloById)

	app.Post("GOairways/API/aviones/create", func(c *fiber.Ctx) error {
		return c.SendString("Aqui se creara un avion")
	})

	app.Get("GOairways/API/aereopuertos", func(c *fiber.Ctx) error {
		return c.SendString("Aqui se listaran los aereopuertos")
	})
}
