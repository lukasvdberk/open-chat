package direct

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/opensource-discord/auth"
	"github.com/lukasvdberk/opensource-discord/config"
	"github.com/lukasvdberk/opensource-discord/friend"
	"github.com/lukasvdberk/opensource-discord/responses"
	web_notification "github.com/lukasvdberk/opensource-discord/web-notification"
	"strconv"
)

func GetRoutes(app *fiber.App) *fiber.App {
	// This the channel the server will set for direct-message notifications.
	messageChannel := "direct-messages"
	// All the routes for receiving for fetching and messages with friend (dms's not servers)

	// Add message to conversation
	app.Post(config.GetDefaultApiRoute()+"/messages", func(c *fiber.Ctx) {
		// POST data required:
		// 		friendUserId: id of the friend you want to message
		//		messageContent: content as a string you wish to send.
		postData := make(map[string]interface{})
		err := json.Unmarshal([]byte(c.Body()), &postData)

		// TODO this parsing could probably be done more cleanly
		var friendUserId int64 = 0
		var tmpVar float64
		var parsed bool
		tmpVar, parsed = postData["friendUserId"].(float64)

		if parsed {
			friendUserId = int64(int(tmpVar))
		}

		if err == nil && friendUserId != 0 {
			jwtClaim := auth.GetJWTClaimsFromContext(c)

			if jwtClaim != nil && friendUserId != 0 {
				friendRelationId := friend.GetFriendRelation(jwtClaim.Id, friendUserId)

				// -1 means not found. else it means they are not friend
				if friendRelationId != -1 {
					message := new(FriendMessage)

					friendRelation := new(friend.Friend)
					friendRelation.Id = friendRelationId

					message.FriendRelation = *friendRelation
					message.FromUser = jwtClaim.Id
					message.ReadMessage = false
					message.MessageContent = postData["messageContent"].(string)

					message, err = SaveMessage(message)

					if err == nil {
						// Send push notification to other user

						// TODO Fetch the right data
						popUpData := web_notification.NotificationPopUpData{
							Title:          friendRelation.User1.Username,
							MessageContent: message.MessageContent,
							Icon:           friendRelation.User1.ProfilePhoto,
						}

						web_notification.PushNotificationToUser(friendUserId, web_notification.Notification{
							Data:                  message,
							MessageChannel:        messageChannel,
							NotificationPopUpData: popUpData,
						})

						responses.SuccessResponse(fiber.Map{
							"message":        "successfully saved message",
							"messageContent": message,
						}, c)
					} else {
						responses.ErrorResponse(4, fiber.Map{
							"errorMessage": "failed to save message",
						}, c)
					}
				} else {
					responses.ErrorResponse(3, fiber.Map{
						"errorMessage": "User is a not a friend of the submitted user id",
					}, c)
				}
			} else {
				responses.ErrorResponse(2, fiber.Map{
					"errorMessage": "failed to retrieve one of the ids",
				}, c)
			}
		} else {
			responses.ErrorResponse(1, fiber.Map{
				"errorMessage": "not a valid user id",
			}, c)
		}
	})

	// get messages for a conversation
	app.Get(config.GetDefaultApiRoute()+"/messages/:friendId", func(c *fiber.Ctx) {
		friendIdStr := c.Params("friendId")
		if _, err := strconv.Atoi(friendIdStr); err == nil {
			friendId, _ := strconv.ParseInt(friendIdStr, 10, 64)

			friendRelationId := friend.GetFriendRelation(friendId, auth.GetJWTClaimsFromContext(c).Id)

			// -1 means not found
			if friendRelationId != -1 {
				responses.SuccessResponse(fiber.Map{
					"messages": GetMessagesFromFriend(friendRelationId),
				}, c)
			} else {
				responses.ErrorResponse(1, fiber.Map{
					"errorMessage": "friend id not found",
				}, c)
			}
		} else {
			responses.ErrorResponse(2, fiber.Map{
				"errorMessage": "not a valid friend id",
			}, c)
		}
	})

	return app
}
