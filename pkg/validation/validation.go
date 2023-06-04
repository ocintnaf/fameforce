package validation

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ocintnaf/fameforce/pkg/helpers"
)

func Validate(data any) []string {
	validate := validator.New()

	translator := helpers.EnglishTranslator()
	_ = en_translations.RegisterDefaultTranslations(validate, translator)

	err := validate.Struct(data)
	errs := translateError(err, translator)

	return errs
}

func translateError(err error, trans ut.Translator) []string {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	var errs []string

	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr.Error())
	}

	return errs
}
