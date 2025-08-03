package handler

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/rakhiazfa/fiber-boilerplate/internal/constants"
	"github.com/rakhiazfa/fiber-boilerplate/internal/service"
)

type HealthcheckHandler struct {
	healthcheckService *service.HealthcheckService
}

func NewHealthcheckHandler(healthcheckService *service.HealthcheckService) *HealthcheckHandler {
	return &HealthcheckHandler{healthcheckService}
}

func (h *HealthcheckHandler) Check(c fiber.Ctx) error {
	err := h.healthcheckService.Check()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status":    constants.HealthStatusUnhealthy,
			"message":   err.Error(),
			"timestamp": time.Now(),
		})
	}

	return c.JSON(fiber.Map{
		"status":    constants.HealthStatusHealthy,
		"timestamp": time.Now(),
	})
}
