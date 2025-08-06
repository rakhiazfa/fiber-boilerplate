package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/handler"
)

type AuthRouter struct {
	handler *handler.AuthHandler
}

func NewAuthRouter(handler *handler.AuthHandler) *AuthRouter {
	return &AuthRouter{handler}
}

func (r *AuthRouter) Load(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Post("/sign-up", r.handler.SignUp)
}
