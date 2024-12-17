package controllers

import (
	"aereopuertos/models"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("HOLA DESDE AEREOPUERTOS 🛫")
}

func ObtenerAereopuertos(c *fiber.Ctx) error {
	return c.JSON(models.GetAllAirports())
}
