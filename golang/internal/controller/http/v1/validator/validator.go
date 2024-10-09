package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	e "github.com/nikitaSstepanov/tools/error"
)

func Struct(s interface{}, args ...Arg) e.Error {
	validate := validator.New()
	setupArgs(validate, args)

	err := validate.Struct(s)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		msg := fmt.Sprintf("Incorrect data: %s", errors)
		return e.New(msg, e.BadInput)
	}

	return nil
}

func StringLength(s string, min int, max int) e.Error {
	if len(s) < min || len(s) > max {
		return lenErr
	}

	return nil
}
