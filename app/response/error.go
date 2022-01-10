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
	switch err.(type) {
	case validator.ValidationErrors:
		m := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			m[err.Field()] = err.Tag()
		}
		return NewErrorResponse(m)
	default:
		return NewBadRequestResponse()
	}
}
