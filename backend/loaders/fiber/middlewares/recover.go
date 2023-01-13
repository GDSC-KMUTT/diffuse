package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"backend/utils/config"
)

var Recover = func() fiber.Handler {
	if config.C.Environment == 1 {
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	return recover.New()
}()
