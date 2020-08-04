package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/lukasvdberk/opensource-discord/responses"
	"time"
)

func GetRoutes(app *fiber.App) *fiber.App {
	// TODO refactor with base api route
	// TODO add proper struct with respond codes or something
	app.Post("/api/auth/login", func(c *fiber.Ctx) {
		c.Accepts("application/json")
		user := new(User)

		_ = c.BodyParser(user)

		if user.Password != "" && user.Username != "" {

			isValidUser := CheckUserCredentials(&user)

			if isValidUser {
				// Create JWT token
				token := jwt.New(jwt.SigningMethodHS256)
				claims := token.Claims.(jwt.MapClaims)
				claims["userId"] = user.Id
				// Good for a month
				claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

				// Generate encoded token and send it as response.
				t, err := token.SignedString([]byte(getSecret()))
				if err != nil {
					// TODO add custom error handler
					c.SendStatus(fiber.StatusInternalServerError)
					return
				}

				responses.Response(fiber.Map{
					"token": t,
				}, c)
			} else {
				responses.Response(fiber.Map{
					"code":    2,
					"message": "User with specified information does not exist",
				}, c)
			}
		} else {
			invalidInformationResponse(c)
		}
	})

	app.Post("/api/auth/register", func(c *fiber.Ctx) {
		c.Accepts("application/json")
		user := new(User)

		_ = c.BodyParser(user)

		if user.Password != "" && user.Username != "" {
			userWithId := RegisterUser(user)

			if userWithId != nil {
				successResponse("Successfully registered user.", c)
			}
		} else {
			invalidInformationResponse(c)
		}
	})

	app.Get("/api/auth/get-new-token", func(c *fiber.Ctx) {
		c.Accepts("application/json")
		c.Send("yes")
	})

	return app
}

func invalidInformationResponse(c *fiber.Ctx) {
	responses.Response(fiber.Map{
		"code":    1,
		"message": "Did not send username and password",
	}, c)
}

func successResponse(message string, c *fiber.Ctx) {
	responses.Response(fiber.Map{
		"code":    0,
		"message": message,
	}, c)
}

func getSecret() string {
	// TODO replace with .env variable
	// TODO put this function in some sort config thing
	return "ceb020a5-7a88-498d-a5a8-7e30d6dbb524"
}
