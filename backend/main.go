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

//func login(c *fiber.Ctx) {
//	user := c.FormValue("user")
//	pass := c.FormValue("pass")
//
//	// Throws Unauthorized error
//	if user != "john" || pass != "doe" {
//		c.SendStatus(fiber.StatusUnauthorized)
//		return
//	}
//
//	// Create token
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	// Set claims
//	claims := token.Claims.(jwt.MapClaims)
//	claims["name"] = "Lukas van den Berk"
//	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//
//	// Generate encoded token and send it as response.
//	t, err := token.SignedString([]byte(getSecret()))
//	if err != nil {
//		c.SendStatus(fiber.StatusInternalServerError)
//		return
//	}
//
//	c.JSON(fiber.Map{"token": t})
//}
//
//func accessible(c *fiber.Ctx) {
//	c.Send("Accessible")
//}
//
//func restricted(c *fiber.Ctx) {
//	user := c.Locals("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//	name := claims["name"].(string)
//	c.Send("Welcome " + name)
//}
