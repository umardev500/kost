package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(id, u string, dur time.Duration) (t string, err error) {

	secret := os.Getenv("SECRET")
	claims := jwt.MapClaims{
		"id":       id,
		"username": u,
		"exp":      time.Now().UTC().Add(dur).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err = token.SignedString([]byte(secret))

	return
}

func GetMapclaims(c *fiber.Ctx) jwt.MapClaims {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return claims
}
