package direct

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/opensource-discord/auth"
	"github.com/lukasvdberk/opensource-discord/config"
	"github.com/lukasvdberk/opensource-discord/friend"
	"github.com/lukasvdberk/opensource-discord/responses"
)

func GetRoutes(app *fiber.App) *fiber.App {
	// All the routes for receiving for fetching and messages with friend (dms's not servers)

	// Add message to conversation
	app.Post(config.GetDefaultApiRoute()+"/messages", func(c *fiber.Ctx) {
		// POST data required:
		// 		friendUserId: id of the friend you want to message
		//		messageContent: content as a string you wish to send.
		postData := make(map[string]interface{})
		err := json.Unmarshal([]byte(c.Body()), &postData)
		var id, errParsingId = postData["friendUserId"].(int64)

		if err == nil && !errParsingId {
			friendUserId := id
			jwtClaim := auth.GetJWTClaimsFromContext(c)

			if jwtClaim != nil {
				// TODO this currently will always return true. Actually implement the function.
				if friend.IsUserFriend(jwtClaim.Id, friendUserId) {
					message := new(FriendMessage)

					fromUser := new(auth.User)
					fromUser.Id = jwtClaim.Id

					friendRelation := new(friend.Friend)

					// TODO actually set the right id.
					friendRelation.Id = 2

					message.FriendRelation = *friendRelation
					message.FromUser = *fromUser
					message.ReadMessage = false
					message.MessageContent = postData["messageContent"].(string)

					_, err := SaveMessage(*message)

					if err == nil {
						responses.SuccessResponse(fiber.Map{
							"message": "successfully saved message",
						}, c)
					} else {
						responses.ErrorResponse(1, fiber.Map{
							"errorMessage": "failed to save message",
						}, c)
					}

					// TODO send message
				}
			}
			// add message with timestamp.
		} else {
			responses.ErrorResponse(1, fiber.Map{
				"errorMessage": "not a valid user id",
			}, c)
		}
	})

	return app
}
