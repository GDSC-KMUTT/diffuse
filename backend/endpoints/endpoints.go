package endpoints

import (
	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("/account")
	_ = account
}
