package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gosuit/e"
)

func Struct(s any, args ...Arg) e.Error {
	validate := validator.New()
	if err := setupArgs(validate, args); err != nil {
		return e.New("Incorrect data", e.BadInput, err)
	}

	err := validate.Struct(s)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		msg := ""

		for _, err := range errors {
			switch err.Tag() {
			case "required":
				msg += "Field" + err.Field() + "is required. "
			case "email":
				msg += "Invalid email. "
			case "min":
				msg += "Min length of " + err.Field() + " is " + err.Param() + ". "
			case "max":
				msg += "Max length of " + err.Field() + " is " + err.Param() + ". "
			case "password":
				msg += "Password must include latin letters in upper and lower case, numbers and special symbols. "
			case "age":
				msg += "Minimal avaliable age is 8. "
			case "fields":
				msg += "Invalid fields. "
			}
		}

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
