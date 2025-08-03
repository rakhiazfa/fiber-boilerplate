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

var healthcheckModule = wire.NewSet(
	service.NewHealthcheckService,
	handler.NewHealthcheckHandler,
	router.NewHealthcheckRouter,
)

func NewApplication() *fiber.App {
	wire.Build(
		database.NewPostgreSQLConnection,
		handler.NewErrorHandler,
		healthcheckModule,
		router.New,
	)

	return nil
}
