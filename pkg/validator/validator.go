package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	gpv "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3/log"
)

var ErrorMessages = map[string]string{
	"required": "/f is required",
	"email":    "/f must be a valid email address",
	"username": "/f must be a valid username",
	"numeric":  "/f must be a number",
	"alphanum": "/f must contain alphanumeric characters",
	"boolean":  "/f must be a boolean value",
	"uuid":     "/f must be a valid UUID",
	"min":      "/f must have at least /p characters",
	"max":      "/f must not exceed /p characters",
	"gte":      "/f must be at least /p",
	"lte":      "/f must be at most /p",
	"eqfield":  "/f must be the same as the /p field",
}

type Validator struct {
	validator *gpv.Validate
}

func New() *Validator {
	validator := gpv.New()

	err := validator.RegisterValidation("username", usernameValidation)
	if err != nil {
		log.Panic(err)
	}

	return &Validator{validator}
}

func (v *Validator) Validate(i any) error {
	return v.validator.Struct(i)
}

func FormatValidationErrors(errors validator.ValidationErrors) map[string]string {
	formattedErrors := make(map[string]string)

	for _, err := range errors {
		field := err.Field()
		tag := err.Tag()
		param := err.Param()

		msgTemplate := ErrorMessages[tag]

		if msgTemplate == "" {
			msgTemplate = fmt.Sprintf("%s validation failed on %s", field, tag)
		}

		msg := strings.ReplaceAll(msgTemplate, "/f", field)
		msg = strings.ReplaceAll(msg, "/p", param)

		formattedErrors[field] = msg
	}

	return formattedErrors
}
