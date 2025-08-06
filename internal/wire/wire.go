//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/wire"
	"github.com/rakhiazfa/fiber-boilerplate/internal/config/application"
	"github.com/rakhiazfa/fiber-boilerplate/internal/config/database"
	"github.com/rakhiazfa/fiber-boilerplate/internal/config/logger"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/handler"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/router"
	"github.com/rakhiazfa/fiber-boilerplate/internal/repository"
	"github.com/rakhiazfa/fiber-boilerplate/internal/service"
)

var healthCheckModule = wire.NewSet(
	service.NewHealthCheckService,
	handler.NewHealthCheckHandler,
	router.NewHealthCheckRouter,
)

var userModule = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
)

var authModule = wire.NewSet(
	service.NewAuthService,
	handler.NewAuthHandler,
	router.NewAuthRouter,
)

func Bootstrap() *fiber.App {
	wire.Build(
		logger.New,
		database.NewPostgreSQLConnection,
		handler.NewErrorHandler,
		healthCheckModule,
		userModule,
		authModule,
		application.New,
	)

	return nil
}
