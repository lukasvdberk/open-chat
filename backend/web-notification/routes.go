package web_notification

import (
	"encoding/json"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/opensource-discord/auth"
	"github.com/lukasvdberk/opensource-discord/config"
	"github.com/lukasvdberk/opensource-discord/responses"
	"os"
)

func GetRoutes(app *fiber.App) *fiber.App {
	// Registration for web notification on the web
	// Add message to conversation
	app.Post(config.GetDefaultApiRoute()+"/web-notifications/add-device", func(c *fiber.Ctx) {
		parsedDevice := &webpush.Subscription{}
		err := json.Unmarshal([]byte(c.Body()), &parsedDevice)

		if err != nil {
			responses.ErrorResponse(1, fiber.Map{
				"errorMessage": "Failed to parse web push subscribe Data",
			}, c)
		}

		jwtClaims := auth.GetJWTClaimsFromContext(c)
		webNotificationDevice := NotificationDevice{
			userId:   jwtClaims.Id,
			endpoint: parsedDevice.Endpoint,
			authKey:  parsedDevice.Keys.Auth,
			p256dh:   parsedDevice.Keys.P256dh,
		}

		isDeviceSaved := SaveNotificationDevice(webNotificationDevice)

		if isDeviceSaved {
			responses.SuccessResponse(fiber.Map{
				"message": "Success! Saved device to database!",
			}, c)
		} else {
			responses.ErrorResponse(2, fiber.Map{
				"errorMessage": "Failed to save web-notification device to server.",
			}, c)
		}
	})

	app.Get(config.GetDefaultApiRoute()+"/web-notifications/get_public_key", func(c *fiber.Ctx) {
		publicKey := os.Getenv("WEB_PUSH_PUBLIC_KEY")

		responses.SuccessResponse(fiber.Map{
			"publicKey": publicKey,
		}, c)
	})

	return app
}
