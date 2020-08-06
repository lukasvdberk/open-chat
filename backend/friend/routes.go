package friend

import (
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/opensource-discord/auth"
	"github.com/lukasvdberk/opensource-discord/config"
	"github.com/lukasvdberk/opensource-discord/responses"
)

func GetRoutes(app *fiber.App) *fiber.App {
	// All the routes for receiving data about friend
	app.Get(config.GetDefaultApiRoute()+"/friend", func(c *fiber.Ctx) {
		// get user id from jwt claim
		// get all friend that exist that have the user id with sql
		// extract that information and remove where the user which is the same as the friend
		jwtClaim := auth.GetJWTClaimsFromContext(c)

		friendsFromUser := GetFriendsByUserId(jwtClaim.Id)

		if friendsFromUser != nil {
			contentMap := fiber.Map{"friends": friendsFromUser}
			responses.SuccessResponse(contentMap, c)
		} else {
			responses.ErrorResponse(1, fiber.Map{
				"errorMessage": "Failed to retrieve friends. 502",
			}, c)
		}
	})
	return app
}
