package auth

import "os"

func getJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
