package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

type JwtClaims struct {
	// user database id
	Id         int64
	Expiration float64
}

func GetJWTClaimsFromContext(c *fiber.Ctx) *JwtClaims {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// The claims are created during the auth process. This is where these key value pair reside
	jwtClaim := new(JwtClaims)
	jwtClaim.Id = int64(claims["userId"].(float64))
	jwtClaim.Expiration = claims["exp"].(float64)

	return jwtClaim
}
