package validations

import (
	"github.com/tariq-ventura/Proyecto-go/internal/logs"

	"github.com/go-playground/validator"
)

var StructValidator = validator.New()

func ValidateStruct(v *validator.Validate, s interface{}) bool {
	err := v.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logs.LogInfo("Validation error", map[string]interface{}{"field": err.Field(), "tag": err.Tag(), "value": err.Value()})
		}
		return false
	} else {
		return true
	}
}
