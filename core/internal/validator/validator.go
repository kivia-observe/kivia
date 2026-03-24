package validator

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)
var (
	validate *validator.Validate
	once     sync.Once
)

func Get() *validator.Validate {
	
	once.Do(func(){
		
		validate = validator.New()
	})
	
	return validate
}

func FirstError(err error) string {
	valErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return "invalid request"
	}

	e := valErrs[0]

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", e.Field())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s failed %s validation", e.Field(), e.Tag())
	}
}