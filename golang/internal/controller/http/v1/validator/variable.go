package validator

import (
	"unicode"

	"github.com/go-playground/validator/v10"
	e "github.com/nikitaSstepanov/tools/error"
)

const (
	Password Arg = iota
)

var (
	lenErr = e.New("Bad string length", e.BadInput)
)

type Arg int

func setupArgs(validate *validator.Validate, args []Arg) {
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case Password:
			validate.RegisterValidation("password", validatePassword)
		default:
			return
		}
	}
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
