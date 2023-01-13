package endpoints

import (
	"github.com/gofiber/fiber/v2"

	homeEndpoints "backend/endpoints/home"
)

func Init(router fiber.Router) {
	// * Account
	account := router.Group("account/")
	_ = account

	// * Home
	home := router.Group("home/")
	home.Get("info", homeEndpoints.InfoHandler)
}
