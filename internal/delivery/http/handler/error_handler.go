package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/sirupsen/logrus"
)

type ErrorHandler struct {
	log *logrus.Logger
}

func NewErrorHandler(log *logrus.Logger) *ErrorHandler {
	return &ErrorHandler{log}
}

func (h *ErrorHandler) Handle(c fiber.Ctx, err error) error {
	statusCode := fiber.StatusInternalServerError

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		statusCode = fiberErr.Code
	}

	statusMessage := http.StatusText(statusCode)

	h.log.WithFields(logrus.Fields{
		"requestId":     requestid.FromContext(c),
		"statusCode":    statusCode,
		"statusMessage": statusMessage,
		"method":        c.Method(),
		"path":          c.Path(),
		"ip":            c.IP(),
		"userAgent":     c.Get("User-Agent"),
	}).Error(err.Error())

	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusMessage,
		"message": err.Error(),
	})
}
