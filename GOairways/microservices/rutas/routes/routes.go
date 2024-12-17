package routes

import (
	"github.com/gofiber/fiber/v2"
	"rutas/models"
	"strconv"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HOLA DESDE RUTAS 📍🗺️")
	})

	app.Get("/rutas", func(c *fiber.Ctx) error {
		return c.JSON(models.GetAllRoutes())
	})

	app.Get("/rutas/active", func(c *fiber.Ctx) error {
		return c.JSON(models.GetActiveRoutes())
	})

	app.Get("/rutas/inactive", func(c *fiber.Ctx) error {
		return c.JSON(models.GetInactiveRoutes())
	})

	app.Get("/rutas/accept/:id", func(c *fiber.Ctx) error {
		number, _ := strconv.Atoi(c.Params("id"))
		return c.JSON(models.ActivateRouteByID(uint(number)))
	})

	app.Get("/rutas/nextID", func(c *fiber.Ctx) error {
		nextID, err := models.GetNextRouteID()
		if err != nil {
			return c.Status(500).SendString("Error al obtener el siguiente ID de ruta")
		}
		return c.JSON(fiber.Map{"next_id": nextID})
	})

	app.Post("/rutas/create", func(c *fiber.Ctx) error {
		route := new(models.Ruta)
		if err := c.BodyParser(route); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		return c.JSON(models.CreateRoute(route))
	})

	app.Post("/rutas/schedule/create", func(c *fiber.Ctx) error {
		schedule := new(models.Horario)
		if err := c.BodyParser(schedule); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		return c.JSON(models.CreateSchedule(schedule))
	})

	app.Post("/rutas/modelRoute/create", func(c *fiber.Ctx) error {
		modelRoute := new(models.ModeloRuta)
		if err := c.BodyParser(modelRoute); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		return c.JSON(models.CreateModelRoute(modelRoute))
	})
}
