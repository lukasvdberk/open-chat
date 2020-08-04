package auth

import (
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/opensource-discord/responses"
)

func GetRoutes(app *fiber.App) *fiber.App {
	// TODO refactor with base api route
	app.Get("/api/auth/login", func(c *fiber.Ctx) {
		c.Send("yes")
	})

	app.Post("/api/auth/register", func(c *fiber.Ctx) {
		c.Accepts("application/json")
		user := new(User)

		_ = c.BodyParser(user)

		if user.Password != "" && user.Username != "" {
			userWithId := RegisterUser(user)

			if userWithId != nil {
				responses.Response(fiber.Map{
					"code":    200,
					"message": "Successfully registered user.",
				}, c)
			}
		} else {
			responses.Response(fiber.Map{
				"code":    400,
				"message": "Did not send username and password",
			}, c)
		}
	})

	app.Get("/api/auth/get-new-token", func(c *fiber.Ctx) {
		c.Accepts("application/json")
		c.Send("yes")
	})

	return app
}
