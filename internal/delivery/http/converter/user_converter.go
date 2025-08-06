package converter

import (
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/dto/request"
	"github.com/rakhiazfa/fiber-boilerplate/internal/entity"
)

func CreateUserRequestToEntity(req *request.CreateUserRequest) *entity.User {
	return &entity.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}
