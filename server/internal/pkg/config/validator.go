package config

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	return validator.New(
		validator.WithRequiredStructEnabled(),
		func(v *validator.Validate) {
			v.RegisterTagNameFunc(func(field reflect.StructField) string {
				return field.Tag.Get("name")
			})
		},
	)
}
