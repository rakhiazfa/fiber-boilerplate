package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rakhiazfa/fiber-boilerplate/internal/constants"
	"github.com/rakhiazfa/fiber-boilerplate/internal/service"
)

type HealthCheckHandler struct {
	healthCheckService *service.HealthCheckService
}

func NewHealthCheckHandler(healthCheckService *service.HealthCheckService) *HealthCheckHandler {
	return &HealthCheckHandler{healthCheckService}
}

func (h *HealthCheckHandler) Check(c fiber.Ctx) error {
	err := h.healthCheckService.Check(c.RequestCtx())
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status":  constants.HealthStatusUnhealthy,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": constants.HealthStatusHealthy,
	})
}
