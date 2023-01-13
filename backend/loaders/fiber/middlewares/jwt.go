package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"

	"backend/types/common"

	"backend/types/response"
	"backend/utils/config"
)

var Jwt = func() fiber.Handler {
	conf := jwtware.Config{
		SigningKey:  []byte(config.C.JwtSecret),
		TokenLookup: "cookie:user",
		ContextKey:  "user",
		Claims:      &common.UserClaims{},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return &response.ErrorInstance{
				Message: "JWT validation failure",
				Err:     err,
			}
		},
	}

	return jwtware.New(conf)
}()
