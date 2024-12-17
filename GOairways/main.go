package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true, // Permitir el envío de cookies
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))

	app.Post("GOairways/API/auth/generatePIN", GenerateNewPin)
	app.Post("GOairways/API/auth/login", LoginFunc)

	//	app.Use(func(c *fiber.Ctx) error {
	//		tokenSting := c.Cookies("jwt")
	//		if tokenSting == "" {
	//			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//				"error": "Unauthorized",
	//			})
	//		}
	//
	//		token, err := jwt.Parse(tokenSting, func(token *jwt.Token) (interface{}, error) {
	//			return []byte("FlySafe"), nil
	//		})
	//
	//		if err != nil || !token.Valid {
	//			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//				"error": "Invalid token",
	//			})
	//		}
	//
	//		claims := token.Claims.(jwt.MapClaims)
	//		c.Locals("userClaims", claims)
	//
	//		return c.Next()
	//	})

	// Routes
	Setup(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("estoEsUnSecretoSobreUnaAereolineaMuyBonitaYmuySegura")},
	}))

	app.Get("GOairways/API/pruebaJWT", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Esto es una prueba de JWT, encontrado",
		})
	})

	app.Listen(":3010")
}
