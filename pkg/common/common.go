package common

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

func EnglishTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	return trans
}
