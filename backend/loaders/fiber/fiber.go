package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"backend/endpoints"

	"backend/loaders/fiber/middlewares"
	"backend/types/response"
	"backend/utils/config"
)

func Init() {
	// Initialize fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler:  ErrorHandler,
		Prefork:       false,
		StrictRouting: true,
		ServerHeader:  config.C.ServerHeader,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	// Init root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.Info("DIFFUSE_ROOT_V1"))
	})

	// Init API endpoints
	apiGroup := app.Group("api/")

	apiGroup.Use(middlewares.Limiter)
	apiGroup.Use(middlewares.Cors)
	apiGroup.Use(middlewares.Recover)

	endpoints.Init(apiGroup)

	// Init not found handler
	app.Use(NotFoundHandler)

	// Startup
	err := app.Listen(config.C.Address)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
