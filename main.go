package main

import (
	"apiv1/src/models"
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) {
		var user models.User

		if err := c.BodyParser(&user); err != nil {
			c.Send(err)
		}
		fmt.Println(user)
		c.Send("Hello, World!")
	})

	app.Listen(3000)
}
