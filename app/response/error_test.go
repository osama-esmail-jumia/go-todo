package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrorResponse(t *testing.T) {
	resp := NewErrorResponse("foo")

	assert.Equal(t, resp.Message, "foo")
}
