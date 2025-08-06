package service

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/converter"
	"github.com/rakhiazfa/fiber-boilerplate/internal/delivery/http/dto/request"
	"github.com/rakhiazfa/fiber-boilerplate/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserService struct {
	log            *logrus.Logger
	db             *gorm.DB
	userRepository *repository.UserRepository
}

func NewUserService(
	log *logrus.Logger,
	db *gorm.DB,
	userRepository *repository.UserRepository,
) *UserService {
	return &UserService{log, db, userRepository}
}

func (s *UserService) Create(ctx context.Context, req *request.CreateUserRequest) error {
	err := s.userRepository.Transaction(ctx, func(tx *gorm.DB) error {
		if err := s.validateUsername(ctx, req.Username); err != nil {
			return err
		}
		if err := s.validateEmail(ctx, req.Email); err != nil {
			return err
		}

		user := converter.CreateUserRequestToEntity(req)

		return s.userRepository.WithTx(tx).Create(user)
	})
	if err != nil {
		s.log.Error("Failed to create user")
		return err
	}

	return nil
}

func (s *UserService) validateUsername(ctx context.Context, username string, excludedIds ...uuid.UUID) error {
	count, err := s.userRepository.WithContext(ctx).CountBy("username", username, excludedIds...)
	if err != nil {
		return err
	}

	if count > 0 {
		return fiber.NewError(fiber.StatusConflict, "User already exists")
	}

	return nil
}

func (s *UserService) validateEmail(ctx context.Context, email string, excludedIds ...uuid.UUID) error {
	count, err := s.userRepository.WithContext(ctx).CountBy("email", email, excludedIds...)
	if err != nil {
		return err
	}

	if count > 0 {
		return fiber.NewError(fiber.StatusConflict, "User already exists")
	}

	return nil
}
