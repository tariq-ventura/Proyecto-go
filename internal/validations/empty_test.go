package validations

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidateStruct(t *testing.T) {
	v := validator.New()

	type TestStruct struct {
		Name string `validate:"required"`
		Age  int    `validate:"gte=18"`
	}

	t.Run("Valid struct", func(t *testing.T) {
		s := TestStruct{Name: "John", Age: 25}
		isValid := ValidateStruct(v, s)
		assert.True(t, isValid)
	})

	t.Run("Invalid struct - missing required field", func(t *testing.T) {
		s := TestStruct{Name: "", Age: 25}
		isValid := ValidateStruct(v, s)
		assert.False(t, isValid)
	})

	t.Run("Invalid struct - value does not meet tag criteria", func(t *testing.T) {
		s := TestStruct{Name: "Jane", Age: 17}
		isValid := ValidateStruct(v, s)
		assert.False(t, isValid)
	})
}
