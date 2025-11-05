package application

import (
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

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

type Application struct {
	Log      *logrus.Logger
	DB       *gorm.DB
	fiberApp *fiber.App
}

func New(
	log *logrus.Logger,
	db *gorm.DB,
	errorHandler *handler.ErrorHandler,
	healthCheckRouter *router.HealthCheckRouter,
) *Application {
	fiberApp := fiber.New(fiber.Config{
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		StructValidator: validator.New(),
		CaseSensitive:   true,
		StrictRouting:   true,
		ErrorHandler:    errorHandler.Handle,
		AppName:         config.Get("APP_NAME"),
	})

	fiberApp.Use(requestid.New())
	fiberApp.Use(logger.New(logger.Config{
		Format: "[PID - ${pid}] [${time}] ${status} - ${method} ${path} ${latency}\n",
	}))
	fiberApp.Use(recoverer.New(recoverer.Config{
		EnableStackTrace: config.Get("APP_ENV") != constants.EnvironmentProduction,
	}))

	api := fiberApp.Group("/api")

	healthCheckRouter.Load(api)

	fiberApp.Use(func(c fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "Resource not found")
	})

	return &Application{
		Log:      log,
		DB:       db,
		fiberApp: fiberApp,
	}
}

func (a *Application) Listen(addr string, config ...fiber.ListenConfig) error {
	return a.fiberApp.Listen(addr, config...)
}
