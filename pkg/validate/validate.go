package validate

import (
	"PetAi/pkg/apperror"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// generate validator instance
var Validator = validator.New(validator.WithRequiredStructEnabled())

func Validate(body interface{}) (apperror.ErrorData, error) {
	err := Validator.Struct(body)
	errorStr := ""
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorStr = errorStr + fmt.Sprintf("Field %s Tag %s", err.Field(), err.Tag())
		}
		return apperror.BadRequestValidation(errorStr), err
	}
	return apperror.ErrorData{}, nil
}
