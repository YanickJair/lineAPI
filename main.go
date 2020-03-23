package main

import (
	"apiv1/src/errors"
	"apiv1/src/models"

	"github.com/go-pg/pg/v9"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	/*
	* @description catch any exception that might be throwed
	 */
	rec := func(c *fiber.Ctx) {
		if r := recover(); r != nil {
			message, _ := r.(pg.Error)
			errorCode := errors.ErrorCode{
				Message: message.Error(),
				Code:    500,
			}
			c.JSON(errorCode)
		}
	}

	app.Post("/", func(c *fiber.Ctx) {
		var user models.User

		if err := c.BodyParser(&user); err != nil {
			c.Send(err)
		}

		defer rec(c)

		user.CreateAccount()
	})

	app.Listen(3000)
}
