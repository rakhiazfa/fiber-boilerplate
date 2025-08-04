package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/rakhiazfa/fiber-boilerplate/internal/wire"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		logrus.Fatalf("Failed to load environment variables: %+v", err)
	}

	runtime.GOMAXPROCS(config.GetInt("MAX_PROCS"))

	local, err := time.LoadLocation(config.Get("APP_TIMEZONE"))
	if err != nil {
		logrus.Fatalf("Failed to load location: %+v", err)
	}

	time.Local = local

	addr := fmt.Sprintf("%s:%d", config.Get("APP_HOST"), config.GetInt("APP_PORT"))

	err = wire.Bootstrap().Listen(addr, fiber.ListenConfig{
		EnablePrefork: config.GetBool("ENABLE_PREFORK"),
	})
	if err != nil {
		logrus.Fatalf("Failed to run application: %+v", err)
	}
}
