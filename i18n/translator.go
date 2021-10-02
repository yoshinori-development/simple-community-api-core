package i18n

import (
	"fmt"

	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
)

var Translator ut.Translator

func Init() error {
	ja := ja.New()
	uni := ut.New(ja, ja)
	var found bool
	if Translator, found = uni.GetTranslator("ja"); !found {
		fmt.Println(found)
		return fmt.Errorf("Translator not found.")
	}
	return nil
}
