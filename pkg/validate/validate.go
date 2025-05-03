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
		return apperror.ErrorData{
			Message:         errorStr,
			CodeDescription: "Validation_ERROR",
			CodeValue:       4000,
		}, err
	}
	return apperror.ErrorData{}, nil
}
