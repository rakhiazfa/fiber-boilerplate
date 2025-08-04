package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/handler"
)

type HealthCheckRouter struct {
	handler *handler.HealthCheckHandler
}

func NewHealthCheckRouter(handler *handler.HealthCheckHandler) *HealthCheckRouter {
	return &HealthCheckRouter{handler}
}

func (r *HealthCheckRouter) Load(api fiber.Router) {
	api.Get("/health", r.handler.Check)
}
