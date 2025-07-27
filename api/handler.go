package api

import "github.com/gofiber/fiber/v2"

func RateLimiter(c *fiber.Ctx) error {
    return c.SendString("Hello from /rate-limiter 🎉")
}