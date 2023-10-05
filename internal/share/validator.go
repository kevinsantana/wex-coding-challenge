package share

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type ValidationResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(o interface{}) []*ValidationResponse {
	var errors []*ValidationResponse

	if err := validate.Struct(o); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ValidationResponse{
				FailedField: err.StructField(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}

	return errors
}
