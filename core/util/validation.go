package util

import (
	"github.com/go-playground/validator"
	"noteapp/core/http"
)

var validatorInstance = validator.New()

func Validate(data any) http.AppError {
	var appError http.AppError
	appError.Errors = map[string]any{}
	err := validatorInstance.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			appError.Errors[err.Field()] = err.Tag()
		}
	}
	return appError
}
