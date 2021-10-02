package router

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/yoshinori-development/simple-community-api-main/i18n"
)

type ErrType string

const (
	ErrTypeValidation = ErrType("validation")
	ErrTypeMessage    = ErrType("message")
)

type ErrorResponse struct {
	Type  ErrType     `json:"type"`
	Error interface{} `json:"error"`
}

func RenderValidationError(err error) ErrorResponse {
	errs := err.(validator.ValidationErrors)
	formattedErrs := make(map[string]string, len(errs))
	for _, err := range errs {
		formattedErrs[err.Field()] = err.Translate(i18n.Translator)
	}

	log.Printf("%+v", formattedErrs)
	res := ErrorResponse{
		Type:  ErrTypeValidation,
		Error: formattedErrs,
	}
	return res
}

func RenderMessageError(err error, message string) ErrorResponse {
	log.Printf("%+v", err)
	res := ErrorResponse{
		Type:  ErrTypeMessage,
		Error: message,
	}
	return res
}
