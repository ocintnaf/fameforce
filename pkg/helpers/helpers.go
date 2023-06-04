package helpers

import (
	"errors"
	"reflect"
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

	if len(splitToken[1]) == 0 {
		return "", errors.New("missing bearer token")
	}

	bearerToken := strings.TrimSpace(splitToken[1])

	return bearerToken, nil
}

type ValidationErrors map[string]string

func Validate(data any) ValidationErrors {
	validate := validator.New()

	// In order to return the field names defined in the json representations of struct
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	translator := EnglishTranslator()
	_ = en_translations.RegisterDefaultTranslations(validate, translator)

	validationErrors := validate.Struct(data)
	if validationErrors == nil {
		return nil
	}

	res := make(ValidationErrors)

	for _, err := range validationErrors.(validator.ValidationErrors) {
		field := err.Field()
		res[field] = err.Translate(translator)
	}

	return res
}
