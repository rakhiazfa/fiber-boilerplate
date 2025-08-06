package application

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/rakhiazfa/fiber-boilerplate/internal/constants"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/handler"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/router"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/config"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/validator"
)

func New(
	errorHandler *handler.ErrorHandler,
	healthCheckRouter *router.HealthCheckRouter,
) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		StructValidator: validator.New(),
		CaseSensitive:   true,
		StrictRouting:   true,
		ErrorHandler:    errorHandler.Handle,
		AppName:         config.Get("APP_NAME"),
	})

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[PID - ${pid}] [${time}] ${status} - ${method} ${path} ${latency}\n",
	}))
	app.Use(recoverer.New(recoverer.Config{
		EnableStackTrace: config.Get("APP_ENV") != constants.EnvironmentProduction,
	}))

	api := app.Group("/api")

	healthCheckRouter.Load(api)

	return app
}
