package validation

import (
	"dans-backend-test/app/model"
	"dans-backend-test/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateUser(request model.UserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
