package validator

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type Arg int

func setupArgs(validate *validator.Validate, args []Arg) error {
	for i := range len(args) {
		switch args[i] {
		case Password:
			return validate.RegisterValidation("password", validatePassword)
		default:
			return nil
		}
	}

	return nil
}

func validatePassword(fl validator.FieldLevel) bool {
	pass := fl.Field().String()

	var number, upper, lower, special bool

	for _, s := range pass {
		switch {
		case unicode.IsNumber(s):
			number = true
		case unicode.IsUpper(s):
			upper = true
		case unicode.IsLower(s):
			lower = true
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			special = true
		default:
			return false
		}
	}

	return number && upper && lower && special && len(pass) >= 8
}
