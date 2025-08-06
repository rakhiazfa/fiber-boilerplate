package service

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type HealthCheckService struct {
	db *gorm.DB
}

func NewHealthCheckService(db *gorm.DB) *HealthCheckService {
	return &HealthCheckService{db}
}

func (s *HealthCheckService) Check(ctx context.Context) error {
	sqlDB, err := s.db.WithContext(ctx).DB()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Failed to get generic database")
	}

	err = sqlDB.Ping()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Database connection failed")
	}

	return nil
}
