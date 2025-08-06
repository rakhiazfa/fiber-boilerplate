package service

import (
	"context"

	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/dto/request"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{userService}
}

func (s *AuthService) SignUp(ctx context.Context, req *request.CreateUserRequest) error {
	return s.userService.Create(ctx, req)
}
