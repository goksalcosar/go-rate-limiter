package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goksalcosar/rate-limiter/internal/limiter"
)

func RateLimitMiddleware(limiter limiter.RateLimiter) fiber.Handler {
    return func(c *fiber.Ctx) error {
        IP := c.IP()

        if !limiter.Allow(IP) {
            return c.Status(fiber.StatusTooManyRequests).
                SendString("Too Many Requests 🧱")
        }

        return c.Next()
    }
}