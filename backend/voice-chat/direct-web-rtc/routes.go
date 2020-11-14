package direct_web_rtc

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/opensource-discord/auth"
	"github.com/lukasvdberk/opensource-discord/config"
	"github.com/lukasvdberk/opensource-discord/friend"
	"github.com/lukasvdberk/opensource-discord/responses"
	web_notification "github.com/lukasvdberk/opensource-discord/web-notification"
)

func GetRoutes(app *fiber.App) *fiber.App {
	// This will set up a web rtc connection for voice chat.

	messageCreateOfferChannel := "direct-web-rtc-messages"

	app.Post(config.GetDefaultApiRoute()+"/voice/send-offer", func(c *fiber.Ctx) {
		// Needs friend-id
		postData := make(map[string]interface{})
		err := json.Unmarshal([]byte(c.Body()), &postData)

		if err == nil {
			jwtClaim := auth.GetJWTClaimsFromContext(c)

			var friendUserId int64 = -1
			tmpVar, parsed := postData["friendUserId"].(float64)

			if parsed {
				friendUserId = int64(int(tmpVar))
			}

			if jwtClaim != nil && friendUserId != -1 {
				friendRelationId := friend.GetFriendRelation(jwtClaim.Id, friendUserId)
				// They are friends we can continue.
				if friendRelationId != -1 {
					// TODO they are friends so send notification to that user. with offer data.
					defer web_notification.PushNotificationToUser(friendUserId, web_notification.Notification{
						Data:           postData["offer"],
						MessageChannel: messageCreateOfferChannel,
						NotificationPopUpData: web_notification.NotificationPopUpData{
							Title:          "",
							MessageContent: "",
							Icon:           "",
						},
					})

					responses.SuccessResponse(fiber.Map{
						"success": "Send offer to other client",
					}, c)
				}
			} else {
				responses.ErrorResponse(1, fiber.Map{
					"errorMessage": "Failed to parse friendUserId. Did you set it correctly?",
				}, c)
			}
		} else {
			responses.ErrorResponse(0, fiber.Map{
				"errorMessage": "You did not supply sufficient information",
			}, c)
		}
	})

	return app
}
