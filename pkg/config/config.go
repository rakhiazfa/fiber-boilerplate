package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func Load(filenames ...string) error {
	if err := godotenv.Load(filenames...); err != nil {
		return err
	}

	return nil
}

func Get(key string, def ...string) string {
	value := os.Getenv(key)

	if len(def) > 0 && value == "" {
		return def[0]
	}

	return value
}

func GetInt(key string, def ...int) int {
	value := Get(key)

	if len(def) > 0 && value == "" {
		return def[0]
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Panicf("failed to convert string to int : %+v", err)
	}

	return intValue
}

func GetDuration(key string, def ...time.Duration) time.Duration {
	value := Get(key)

	if len(def) > 0 && value == "" {
		return def[0]
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Panicf("failed to convert string to int : %+v", err)
	}

	return time.Duration(intValue)
}

func GetBool(key string, def ...bool) bool {
	value := Get(key)

	if len(def) > 0 && value == "" {
		return def[0]
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Panicf("failed to convert string to bool : %+v", err)
	}

	return boolValue
}
