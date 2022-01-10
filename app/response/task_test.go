package response

import (
	"github.com/stretchr/testify/assert"
	"go-todo/app/model"
	"testing"
)

func TestNewTaskResponse(t *testing.T) {
	resp := NewTaskResponse(&model.Task{
		Title:       "foo",
		Description: "bar",
		Completed:   true,
	})

	assert.Equal(t, "foo", resp.Title)
	assert.Equal(t, "bar", resp.Description)
	assert.Equal(t, true, resp.Completed)
}

func TestNewTaskListResponse(t *testing.T) {
	resp := NewTaskListResponse([]*model.Task{
		&model.Task{
			Title:       "foo",
			Description: "bar",
			Completed:   true,
		},
		&model.Task{
			Title:       "fooBaz",
			Description: "barBaz",
			Completed:   false,
		},
	}, 10, 2, 1)

	assert.Equal(t, 2, resp.Limit)
	assert.Equal(t, 1, resp.Offset)
	assert.Equal(t, 10, resp.Total)
	assert.Len(t, resp.Rows, 2)
	assert.Equal(t, "foo", resp.Rows[0].Title)
	assert.Equal(t, "fooBaz", resp.Rows[1].Title)
	assert.Equal(t, "bar", resp.Rows[0].Description)
	assert.Equal(t, "barBaz", resp.Rows[1].Description)
	assert.Equal(t, true, resp.Rows[0].Completed)
	assert.Equal(t, false, resp.Rows[1].Completed)
}
