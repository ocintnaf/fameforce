package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ocintnaf/fameforce/types"
)

func EnglishTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	return trans
}

func GetBearerToken(headerGetter types.HeaderGetter) (string, error) {
	token := headerGetter.Get("Authorization")

	splitToken := strings.SplitAfter(token, "Bearer")
	if len(splitToken) != 2 {
		return "", errors.New("invalid bearer token format")
	}

	bearerToken := strings.TrimSpace(splitToken[1])

	return bearerToken, nil
}

func Validate(data any) []string {
	validate := validator.New()

	translator := EnglishTranslator()
	_ = en_translations.RegisterDefaultTranslations(validate, translator)

	err := validate.Struct(data)
	errs := TranslateError(err, translator)

	return errs
}

func TranslateError(err error, trans ut.Translator) []string {
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
