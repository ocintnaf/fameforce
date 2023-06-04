package helpers

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
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
