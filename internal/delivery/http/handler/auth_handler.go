package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/dto/request"
	"github.com/rakhiazfa/fiber-boilerplate/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) SignUp(c fiber.Ctx) error {
	var req request.CreateUserRequest

	if err := c.Bind().Body(&req); err != nil {
		return err
	}

	if err := h.authService.SignUp(c.RequestCtx(), &req); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Successfully created a new account",
	})
}
