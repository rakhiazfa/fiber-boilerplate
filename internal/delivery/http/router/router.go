package router

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/handler"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/config"
)

func New(
	errorHandler *handler.ErrorHandler,
	healthcheckRouter *HealthcheckRouter,
) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		CaseSensitive: true,
		StrictRouting: true,
		ErrorHandler:  errorHandler.Handle,
		AppName:       config.Get("APP_NAME"),
	})

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		CustomTags: map[string]logger.LogFunc{
			"requestid": func(output logger.Buffer, c fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(requestid.FromContext(c))
			},
		},
		Format: "[PID - ${pid}] [${time}] [${requestid}] ${status} - ${method} ${path} ${latency}\n",
	}))

	api := app.Group("/api")

	healthcheckRouter.Load(api)

	return app
}
