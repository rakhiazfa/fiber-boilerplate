package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/rakhiazfa/fiber-boilerplate/internal/wire"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/config"
)

func main() {
	config.LoadEnv()

	runtime.GOMAXPROCS(config.GetInt("MAX_PROCS"))

	local, err := time.LoadLocation(config.Get("APP_TIMEZONE"))
	if err != nil {
		log.Fatal("failed to load location: ", err)
	}

	time.Local = local

	addr := fmt.Sprintf("%s:%d", config.Get("APP_HOST"), config.GetInt("APP_PORT"))

	err = wire.NewApplication().Listen(addr, fiber.ListenConfig{
		EnablePrefork: config.GetBool("ENABLE_PREFORK"),
	})
	if err != nil {
		log.Fatal("failed to run application")
	}
}
