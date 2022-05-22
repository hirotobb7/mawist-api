package validator

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

// TODO: field 日本語対応
func formatMessage(field, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%sは必須です。", field)
	}

	return ""
}

type Validation struct {
	error error
}

func (validation Validation) GetMessages() []string {
	var messages []string

	if validation.error == nil {
		return messages
	}

	for _, err := range validation.error.(validator.ValidationErrors) {
		messages = append(
			messages,
			formatMessage(err.Field(), err.Tag()),
		)
	}

	return messages
}

func Validate(v interface{}) Validation {
	return Validation{validator.New().Struct(v)}
}
