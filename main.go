package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hunzo/go-auth-jwt/services"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		jwtWrap := services.JwtWrapper{
			SecretKey:       "test",
			Issuer:          "TKO",
			ExpirationHours: 10,
		}

		t, err := jwtWrap.GenToken("tko@nida.ac.th")

		if err != nil {
			return c.Send([]byte(err.Error()))
		}

		return c.JSON(fiber.Map{
			"info":  "test",
			"token": t,
			"link":  "http://localhost:8080/validate?token=" + t,
		})
	})

	app.Get("/validate", func(c *fiber.Ctx) error {
		token := c.Query("token")

		// jwtWrap := services.JwtWrapper{
		// 	SecretKey: "test",
		// 	Issuer:    "test add ",
		// }

		jwtWrap := services.JwtWrapper{}

		claims, err := jwtWrap.ValidateToken(token)

		if err != nil {
			// return c.Send([]byte(err.Error()))
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"info": claims,
		})

	})

	log.Fatal(app.Listen(":8080"))

}
