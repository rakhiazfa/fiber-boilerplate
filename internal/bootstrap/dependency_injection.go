package bootstrap

import (
	"github.com/mazrean/kessoku"
	"github.com/rakhiazfa/fiber-boilerplate/internal/bootstrap/application"
	"github.com/rakhiazfa/fiber-boilerplate/internal/bootstrap/database"
	"github.com/rakhiazfa/fiber-boilerplate/internal/bootstrap/logger"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/handler"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/router"
	"github.com/rakhiazfa/fiber-boilerplate/internal/service"
)

var _ = kessoku.Inject[*application.Application]("Init",
	kessoku.Provide(logger.New),
	kessoku.Provide(database.NewPostgreSQLConnection),
	kessoku.Provide(handler.NewErrorHandler),

	kessoku.Set(
		kessoku.Provide(service.NewHealthCheckService),
		kessoku.Provide(handler.NewHealthCheckHandler),
		kessoku.Provide(router.NewHealthCheckRouter),
	),

	kessoku.Provide(application.New),
)
