package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/rakhiazfa/fiber-boilerplate/internal/bootstrap"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.Load(); err != nil {
		logrus.Fatalf("Failed to load environment variables: %+v", err)
	}

	runtime.GOMAXPROCS(config.GetInt("MAX_PROCS"))

	app := bootstrap.Init()

	local, err := time.LoadLocation(config.Get("APP_TIMEZONE"))
	if err != nil {
		app.Log.Fatalf("Failed to load location: %+v", err)
	}

	time.Local = local

	addr := fmt.Sprintf("%s:%d", config.Get("APP_HOST", "127.0.0.1"), config.GetInt("APP_PORT", 8080))

	err = app.Listen(addr, fiber.ListenConfig{
		EnablePrefork: config.GetBool("ENABLE_PREFORK", false),
	})
	if err != nil {
		app.Log.Fatalf("Failed to start application: %+v", err)
	}
}
