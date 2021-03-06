package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

const (
	BAD_REQUEST_MESSAGE     = "Bad request"
	NOT_FOUND_MESSAGE       = "Not found"
	DUPLICATE_TITLE_MESSAGE = "Title should be unique"
)

type Error struct {
	Error interface{} `json:"error"`
}

func NewErrorResponse(err interface{}) *Error {
	return &Error{
		Error: err,
	}
}

func NewNotFoundResponse() *Error {
	return NewErrorResponse(NOT_FOUND_MESSAGE)
}

func NewBadRequestResponse() *Error {
	return NewErrorResponse(BAD_REQUEST_MESSAGE)
}

func NewValidationErrorResponse(err error) *Error {
	if errors, ok := err.(validator.ValidationErrors); ok {
		m := make(map[string]string)
		for _, err := range errors {
			m[err.Field()] = strings.Trim(fmt.Sprintf("%s %s", err.Tag(), err.Param()), " ")
		}
		return NewErrorResponse(m)
	}

	return NewBadRequestResponse()
}
