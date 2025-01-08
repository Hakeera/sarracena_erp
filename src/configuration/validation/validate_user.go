// Package validation provides functions and utilities for input validation and error handling.
package validation

import (
	"encoding/json"
	"errors"
	"sarracena_erp/src/configuration/rest_err"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	// Validate is the global validator instance for input validation.
	Validate = validator.New()
	
	// transl is the global translator for error messages.
	transl ut.Translator
)

// init initializes the validator engine and registers translations.
func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

// ValidateUserError processes and returns a structured error for validation failures.
// validationErr: The validation error to be processed.
// Returns a RestErr object containing the validation details.
func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	errorsCauses := []rest_err.Causes{}

	// Handle JSON Unmarshal type errors.
	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type", "invalid field type", 400, nil)
	}

	// Handle validation errors from the validator package.
	if errors.As(validationErr, &jsonValidationError) {
		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Invalid fields", "validation error", 400, errorsCauses)
	}

	// Handle generic errors.
	return rest_err.NewBadRequestError("Generic case", "Bad Request Error", 400, errorsCauses)
}
