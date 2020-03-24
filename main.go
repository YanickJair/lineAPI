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
			status := errors.ErrorCode{
				Message: message.Error(),
				Code:    500,
			}
			response := models.Response{Status: status, Data: nil}
			c.JSON(response)
		}
	}

	app.Post("/signup", func(c *fiber.Ctx) {
		var user models.User

		if err := c.BodyParser(&user); err != nil {
			c.Send(err)
		}

		defer rec(c)

		status := user.CreateAccount()
		response := models.Response{Status: status, Data: nil}
		c.JSON(response)
	})

	app.Post("/login", func(c *fiber.Ctx) {
		defer rec(c)

		var user models.User
		err := c.BodyParser(&user)

		if err != nil {
			panic(err)
		}

		if len(user.Email) > 0 || len(user.Password) > 0 {
			user, err := models.SignIn(user.Email, user.Password)
			response := models.Response{Data: user, Status: err}
			c.JSON(response)
		} else {
			status := errors.ErrorCode{Message: "Provide email and password", Code: 400}
			response := models.Response{Status: status, Data: nil}
			c.JSON(response)
		}
	})

	app.Listen(3000)
}
