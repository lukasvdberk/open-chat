package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("hello world")
	})

	app.Static("/", "./front-end-public")

	_ = app.Listen("0.0.0.0:4000")
}
