package main

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/lukasvdberk/opensource-discord/auth"
)

func getSecret() string {
	// TODO replace with .env variable
	return "ceb020a5-7a88-498d-a5a8-7e30d6dbb524"
}

func main() {
	app := fiber.New()

	// Authentication
	app = auth.GetRoutes(app)

	// NOTE IMPORTANT: All public routes should be imported before jwt middleware.
	// all routes that requires authentication should be imported after.
	//// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(getSecret()),
	}))

	app.Static("/", "./front-end-public")

	_ = app.Listen("0.0.0.0:4000")
}
