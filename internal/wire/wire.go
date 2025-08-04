//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/wire"
	"github.com/rakhiazfa/fiber-boilerplate/internal/database"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/handler"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/router"
	"github.com/rakhiazfa/fiber-boilerplate/internal/service"
)

var healthCheckModule = wire.NewSet(
	service.NewHealthCheckService,
	handler.NewHealthCheckHandler,
	router.NewHealthCheckRouter,
)

func NewApplication() *fiber.App {
	wire.Build(
		database.NewPostgreSQLConnection,
		handler.NewErrorHandler,
		healthCheckModule,
		router.New,
	)

	return nil
}
