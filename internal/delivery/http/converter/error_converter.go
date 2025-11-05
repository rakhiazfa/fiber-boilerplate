package converter

import (
	"fmt"
	"strings"

	gpv "github.com/go-playground/validator/v10"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/formatter"
	"github.com/rakhiazfa/fiber-boilerplate/pkg/validator"
)

func ConvertValidationErrorsToMap(errors gpv.ValidationErrors) map[string]string {
	formattedErrors := make(map[string]string)

	for _, err := range errors {
		field := formatter.LowerCaseFirst(err.Field())
		tag := err.Tag()
		param := err.Param()

		if tag == "eqfield" {
			param = formatter.LowerCaseFirst(param)
		}

		msgTemplate := validator.ErrorMessages[tag]
		if msgTemplate == "" {
			msgTemplate = fmt.Sprintf("%s validation failed on %s", field, tag)
		}

		msg := strings.ReplaceAll(msgTemplate, "/f", field)
		msg = strings.ReplaceAll(msg, "/p", param)

		formattedErrors[field] = msg
	}

	return formattedErrors
}
