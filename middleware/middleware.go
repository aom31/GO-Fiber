package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CheckMiddlewareLogURL(c *fiber.Ctx) error {

	startTime := time.Now()

	fmt.Printf(
		"URL = %s, Method = %s,Time = %s \n", c.OriginalURL(), c.Method(), startTime)
	return c.Next()
}

func ValidateRoleAuthorize(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
