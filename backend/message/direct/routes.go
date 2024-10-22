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
	// This the channel the server will set for direct-web-rtc-message notifications.
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

						userData := auth.GetUserById(jwtClaim.Id)

						popUpData := web_notification.NotificationPopUpData{
							Title:          userData.Username,
							MessageContent: message.MessageContent,
							Icon:           userData.ProfilePhoto,
						}

						// TODO maybe execute this in a background job
						// just wait with sending the message
						defer web_notification.PushNotificationToUser(friendUserId, web_notification.Notification{
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
	app.Get(config.GetDefaultApiRoute()+"/messages/:friendId/:fromTimestamp", func(c *fiber.Ctx) {

		friendIdStr := c.Params("friendId")

		// is a optional parameter
		lastMessageTimestamp := c.Params("fromTimestamp")

		if _, err := strconv.Atoi(friendIdStr); err == nil {
			friendId, _ := strconv.ParseInt(friendIdStr, 10, 64)

			friendRelationId := friend.GetFriendRelation(friendId, auth.GetJWTClaimsFromContext(c).Id)

			// -1 means not found
			if friendRelationId != -1 {

				if _, err := strconv.Atoi(lastMessageTimestamp); err == nil {
					// With this timestamp we can fetch messages starting from that timestamp.
					lastMessageTimestamp, _ := strconv.ParseInt(lastMessageTimestamp, 10, 64)
					responses.SuccessResponse(fiber.Map{
						"messages": GetMessagesFromFriend(friendRelationId, lastMessageTimestamp),
					}, c)
				} else {
					responses.SuccessResponse(fiber.Map{
						"messages": GetMessagesFromFriend(friendRelationId, -1),
					}, c)
				}
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

	app.Get(config.GetDefaultApiRoute()+"/amount_of_new_messages/", func(c *fiber.Ctx) {
		// Returns a dict of how many new messages a user has for each friend id.
		userId := auth.GetJWTClaimsFromContext(c).Id

		// TODO actually give back a list
		responses.SuccessResponse(fiber.Map{
			"amountOfMessagesPerUserId": GetAmountOfNewMessagesFromUser(userId),
		}, c)
	})

	app.Get(config.GetDefaultApiRoute()+"/read_messages/:friendId", func(c *fiber.Ctx) {
		// Returns a dict of how many new messages a user has for each friend id.
		userId := auth.GetJWTClaimsFromContext(c).Id
		friendIdStr := c.Params("friendId")
		if _, err := strconv.Atoi(friendIdStr); err == nil {
			friendId, _ := strconv.ParseInt(friendIdStr, 10, 64)

			friendRelationId := friend.GetFriendRelation(userId, friendId)

			// -1 means not found
			if friendRelationId != -1 {
				success := ReadMessageAllMessages(friendRelationId)

				if success {
					responses.SuccessResponse(fiber.Map{
						"message": "Updated read message status!",
					}, c)
				} else {
					responses.ErrorResponse(3, fiber.Map{
						"errorMessage": "failed to update database",
					}, c)
				}
			} else {
				responses.ErrorResponse(2, fiber.Map{
					"errorMessage": "Did not find friend relation. Seems they are not friends.",
				}, c)
			}
		} else {
			responses.ErrorResponse(1, fiber.Map{
				"errorMessage": "not a valid friend id",
			}, c)
		}
	})

	return app
}
