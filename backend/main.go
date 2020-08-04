package main

import (
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/opensource-discord/auth"
)

func main() {
	app := fiber.New()

	// NOTE IMPORTANT: All public routes should be imported before auth.GetRoutes.
	// all routes that requires authentication should be imported after.
	app = auth.GetRoutes(app)

	app.Static("/", "./front-end-public")

	_ = app.Listen("0.0.0.0:4000")
}
