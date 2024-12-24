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
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	// Declarar errorsCauses no escopo geral da função
	errorsCauses := []rest_err.Causes{}

	// Caso de erro de tipo de campo (UnmarshalTypeError)
	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Tipo de Campo Inválido", "invalid field type", 400, nil)
	}

	// Caso de erro de validação do Validator
	if errors.As(validationErr, &jsonValidationError) {
		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Campos Inválidos", "validation error", 400, errorsCauses)
	}

	// Caso genérico
	return rest_err.NewBadRequestError("Caso Genérico", "Bad Request Error", 400, errorsCauses)
}
