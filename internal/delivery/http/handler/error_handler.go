package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (h *ErrorHandler) Handle(c fiber.Ctx, err error) error {
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"error":   http.StatusText(fiberErr.Code),
			"message": fiberErr.Message,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   "Internal Server Error",
		"message": err.Error(),
	})
}
