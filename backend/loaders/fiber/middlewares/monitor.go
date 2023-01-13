package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var Monitor = func() fiber.Handler {
	return monitor.New()
}()
