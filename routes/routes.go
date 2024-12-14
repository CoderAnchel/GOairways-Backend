package routes

import (
	"github.com/gofiber/fiber/v2"
	"users/models"
	"users/services"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HOLA DESDE USERS 👨🏻‍💻")
	})

	app.Post("/users/create", func(c *fiber.Ctx) error {
		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Cannot parse JSON!",
			})
		}

		hashedPassword, err := services.HashPassword(user.Password)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Cannot hash password!",
			})
		}

		user.Password = hashedPassword

		if err := models.CreateUser(user); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Cannot create user!",
			})
		}

		return c.JSON("User created successfully!")
	})

	app.Post("/users/login", func(c *fiber.Ctx) error {
		Login := new(models.Login)
		if err := c.BodyParser(Login); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Cannot parse JSON!",
			})
		}

		response, err := models.LoginFunc(Login.DNI, Login.Password, Login.PIN)

		if response == nil || err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Cannot login user!",
			})
		}

		return c.JSON(fiber.Map{
			"user": response,
		})
	})

	app.Post("/users/modifyPIN", func(c *fiber.Ctx) error {
		login := new(models.Login)
		if err := c.BodyParser(login); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Cannot parse JSON!",
			})
		}

		if err := models.ModyfyPIN(login); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Cannot modify PIN!",
			})
		}

		return c.JSON("PIN modified successfully!")
	})
}
