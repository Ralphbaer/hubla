package http

import (
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

func fields(errors validator.ValidationErrors, trans ut.Translator) FieldValidations {
	l := len(errors)
	if l > 0 {

		fields := make(FieldValidations, l)
		for _, e := range errors {
			fields[e.Field()] = e.Translate(trans)
		}

		return fields
	}
	return nil
}

// ValidateStruct validates struct based on their tags
func ValidateStruct(s interface{}) error {
	v, trans := newValidator()
	err := v.Struct(s)
	if err != nil {
		errPtr := malformedRequestErr(err.(validator.ValidationErrors), trans)
		return &errPtr
	}
	return nil
}

func malformedRequestErr(err validator.ValidationErrors, trans ut.Translator) ValidationError {
	return ValidationError{
		Code:    400,
		Message: "Malformed request.",
		Fields:  fields(err, trans),
	}
}

// Validator returns an instance of "gopkg.in/go-playground/validator.v9"
func newValidator() (*validator.Validate, ut.Translator) { return nil, nil }
