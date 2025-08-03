package service

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type HealthcheckService struct {
	db *gorm.DB
}

func NewHealthcheckService(db *gorm.DB) *HealthcheckService {
	return &HealthcheckService{db}
}

func (s *HealthcheckService) Check() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Failed to get generic database")
	}

	err = sqlDB.Ping()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Database connection failed")
	}

	return nil
}
