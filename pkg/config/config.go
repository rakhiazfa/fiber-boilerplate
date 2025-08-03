package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}

func Get(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	value, err := strconv.Atoi(Get(key))
	if err != nil {
		log.Fatal("failed to convert string to int : ", err)
	}

	return value
}

func GetDuration(key string) time.Duration {
	return time.Duration(GetInt(key))
}

func GetBool(key string) bool {
	value, err := strconv.ParseBool(Get(key))
	if err != nil {
		log.Fatal("failed to convert string to bool : ", err)
	}

	return value
}
