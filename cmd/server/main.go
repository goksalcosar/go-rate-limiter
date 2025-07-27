package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goksalcosar/rate-limiter/api"
	"github.com/goksalcosar/rate-limiter/internal/limiter"
	"github.com/goksalcosar/rate-limiter/middleware"
)

func main() {
    app := fiber.New()

    rateLimiter := limiter.NewRateLimiter(limiter.LimiterSlidingWindow)

    app.Use(middleware.RateLimitMiddleware(rateLimiter))

    app.Get("/rate-limiter", api.RateLimiter)

    app.Listen(":3000")
}