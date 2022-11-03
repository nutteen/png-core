package validator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// Validate use a single instance of Validate, it caches struct info
var Validate *validator.Validate

func InitializeValidator()  {
	Validate = validator.New()
	// register function to get tag name from json tags.
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}