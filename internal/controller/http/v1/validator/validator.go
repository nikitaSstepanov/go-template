package validator

import (
	"github.com/go-playground/validator/v10"
	e "github.com/nikitaSstepanov/tools/error"
)

func Struct(s interface{}, args ...Arg) e.Error {
	validate := validator.New()
	if err := setupArgs(validate, args); err != nil {
		return e.New("Incorrect data", e.BadInput, err)
	}

	err := validate.Struct(s)
	if err != nil {
		errors := err.(validator.ValidationErrors)

		return e.New("Incorrect data", e.BadInput, errors)
	}

	return nil
}

func StringLength(s string, min int, max int) e.Error {
	if len(s) < min || len(s) > max {
		return lenErr
	}

	return nil
}
