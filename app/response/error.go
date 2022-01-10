package response

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(msg string) *Error {
	return &Error{
		Message: msg,
	}
}
