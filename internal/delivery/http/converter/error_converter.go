package converter

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/formatter"
)

var errorMessages = map[string]string{
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

func ValidationErrorsToResponse(errors validator.ValidationErrors) map[string]string {
	formattedErrors := make(map[string]string)

	for _, err := range errors {
		field := formatter.LowerCaseFirst(err.Field())
		tag := err.Tag()
		param := err.Param()

		if tag == "eqfield" {
			param = formatter.LowerCaseFirst(param)
		}

		msgTemplate := errorMessages[tag]
		if msgTemplate == "" {
			msgTemplate = fmt.Sprintf("%s validation failed on %s", field, tag)
		}

		msg := strings.Replace(msgTemplate, "/f", field, -1)
		msg = strings.Replace(msg, "/p", param, -1)

		formattedErrors[field] = msg
	}

	return formattedErrors
}
