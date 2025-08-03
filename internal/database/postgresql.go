package database

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3/log"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreSQLConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=%s",
		config.Get("DATABASE_HOST"),
		config.GetInt("DATABASE_PORT"),
		config.Get("DATABASE_NAME"),
		config.Get("DATABASE_USERNAME"),
		config.Get("DATABASE_PASSWORD"),
		config.Get("APP_TIMEZONE"),
	)

	gorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Error("failed to initialize postgres connection: ", err)
	}

	db, err := gorm.DB()
	if err != nil {
		log.Error("failed to get generic database: ", err)
	}

	db.SetMaxOpenConns(config.GetInt("DATABASE_MAX_OPEN_CONNECTIONS"))
	db.SetMaxIdleConns(config.GetInt("DATABASE_MAX_IDLE_CONNECTIONS"))
	db.SetConnMaxLifetime(config.GetDuration("DATABASE_CONNECTION_LIFE_TIME") * time.Minute)
	db.SetConnMaxIdleTime(config.GetDuration("DATABASE_CONNECTION_IDLE_TIME") * time.Minute)

	return gorm
}
