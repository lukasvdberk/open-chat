package main

import (
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/opensource-discord/auth"
	"github.com/lukasvdberk/opensource-discord/friend"
	"github.com/lukasvdberk/opensource-discord/message/direct"
	web_notification "github.com/lukasvdberk/opensource-discord/web-notification"
)

func main() {
	app := fiber.New()

	// NOTE IMPORTANT: All public routes should be imported before auth.GetRoutes.
	// all routes that requires authentication should be imported after.
	app = auth.GetRoutes(app)
	app = friend.GetRoutes(app)

	// direct messages
	app = direct.GetRoutes(app)
	app = web_notification.GetRoutes(app)

	app.Static("/", "./front-end-public")

	_ = app.Listen("0.0.0.0:4000")
}
