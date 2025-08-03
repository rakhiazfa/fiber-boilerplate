package validator

import (
	"regexp"

	gpv "github.com/go-playground/validator/v10"
)

func UsernameValidation(fl gpv.FieldLevel) bool {
	username := fl.Field().String()
	regex := regexp.MustCompile("^[a-zA-Z0-9._]+$")

	return regex.MatchString(username)
}
