package dto

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	etrans "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var translator ut.Translator

func init() {
	validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	translator, _ = uni.GetTranslator("en")
	_ = etrans.RegisterDefaultTranslations(validate, translator)
}

func translateErrors(validationErrors validator.ValidationErrors) map[string]string {
	errorTranslated := make(map[string]string)
	for _, validationError := range validationErrors {
		errorTranslated[validationError.Field()] = validationError.Translate(translator)
	}

	return errorTranslated
}
