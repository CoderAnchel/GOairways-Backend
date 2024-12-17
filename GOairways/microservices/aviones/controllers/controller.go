package controllers

import (
	"aviones/models"
	"aviones/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ObtenerModeloPorTipo(c *fiber.Ctx) error {
	tipo := c.Params("tipo")

	response, err := services.GetModeloByTipo(tipo)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get planes",
		})
	}

	return c.JSON(response)
}

func ObtenerModeloPorID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	response, err := services.GetModeloByID(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get plane",
		})
	}

	return c.JSON(response)
}

func CrearNuevoModelo(c *fiber.Ctx) error {
	modelo := &models.Modelo{}

	if err := c.BodyParser(modelo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := services.CreateNewModelo(modelo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create plane",
		})
	}

	return c.JSON(modelo)
}

func ObtenerModelos(c *fiber.Ctx) error {
	response, err := services.GetAllModelos()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get planes",
		})
	}
	return c.JSON(response)
}

func ObtenerAviones(c *fiber.Ctx) error {
	response := models.GetAllPlanes()

	return c.JSON(response)
}

func ObtenerAvionPorID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	response, err := models.GetPlaneByID(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get plane",
		})
	}

	return c.JSON(response)
}

func CheckMatricula(c *fiber.Ctx) error {
	matricula := c.Params("matricula")

	if models.CheckPlateExists(matricula) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Matricula already exists",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Matricula available",
	})
}
