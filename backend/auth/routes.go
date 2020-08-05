package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/lukasvdberk/opensource-discord/config"
	"github.com/lukasvdberk/opensource-discord/responses"
	"time"
)

func GetRoutes(app *fiber.App) *fiber.App {
	// So we can make request from the browser
	// TODO add a different configuration when deployed or setup nginx.
	app.Use(cors.New())

	app.Post(config.GetDefaultApiRoute()+"/auth/login", func(c *fiber.Ctx) {
		c.Accepts("application/json")
		user := new(User)

		_ = c.BodyParser(user)

		if user.Password != "" && user.Username != "" {
			// Also updates the user id so that is why we need to pass it as a pointer
			isValidUser := CheckUserCredentials(&user)

			if isValidUser {
				token := jwt.New(jwt.SigningMethodHS256)
				claims := token.Claims.(jwt.MapClaims)
				claims["userId"] = user.Id
				// Good for a month
				claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

				// Generate encoded token
				t, err := token.SignedString([]byte(getJWTSecret()))
				if err != nil {
					c.SendStatus(fiber.StatusInternalServerError)
					responses.ErrorResponse(3, fiber.Map{
						"errorMessage": "Failed to create token",
					}, c)
					return
				}

				responses.ErrorResponse(0, fiber.Map{
					"token": t,
				}, c)
			} else {
				responses.ErrorResponse(2, fiber.Map{
					"errorMessage": "User with specified information does not exist",
				}, c)
			}
		} else {
			invalidInformationResponse(c)
		}
	})

	app.Post(config.GetDefaultApiRoute()+"/auth/register", func(c *fiber.Ctx) {
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

	app.Get(config.GetDefaultApiRoute()+"/auth/get-new-token", func(c *fiber.Ctx) {
		c.Accepts("application/json")
		c.Send("yes")
	})

	// setup jwt middleware. all requests after this are authenticated requests.
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(getJWTSecret()),
	}))

	return app
}

func invalidInformationResponse(c *fiber.Ctx) {
	responses.ErrorResponse(1, fiber.Map{
		"errorMessage": "Missing information",
	}, c)
}

func successResponse(message string, c *fiber.Ctx) {
	responses.SuccessResponse(fiber.Map{
		"message": message,
	}, c)
}
