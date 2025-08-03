package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/handler"
)

type HealthcheckRouter struct {
	handler *handler.HealthcheckHandler
}

func NewHealthcheckRouter(handler *handler.HealthcheckHandler) *HealthcheckRouter {
	return &HealthcheckRouter{handler}
}

func (r *HealthcheckRouter) Load(api fiber.Router) {
	api.Get("/health", r.handler.Check)
}
