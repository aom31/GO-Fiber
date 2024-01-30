package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CheckMiddlewareLogURL(c *fiber.Ctx) error {

	startTime := time.Now()

	fmt.Printf(
		"URL = %s, Method = %s,Time = %s \n", c.OriginalURL(), c.Method(), startTime)
	return c.Next()
}
