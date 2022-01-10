package response

import (
	"github.com/go-playground/validator/v10"
)

const (
	BAD_REQUEST_MESSAGE = "Bad request"
)

type Error struct {
	Error interface{} `json:"error"`
}

func NewErrorResponse(err interface{}) *Error {
	return &Error{
		Error: err,
	}
}

func NewBadRequestResponse() *Error {
	return NewErrorResponse(BAD_REQUEST_MESSAGE)
}

func NewValidationErrorResponse(err error) *Error {
	if errors, ok := err.(validator.ValidationErrors); ok {
		m := make(map[string]string)
		for _, err := range errors {
			m[err.Field()] = err.Tag()
		}
		return NewErrorResponse(m)
	}

	return NewBadRequestResponse()
}
