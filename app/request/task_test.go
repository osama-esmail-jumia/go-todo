package request

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_context "go-todo/context/mock"
	"testing"
)

func TestNewTaskListRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := mock_context.NewMockContext(ctrl)
	completed := true
	ctx.EXPECT().BindQuery(&TaskListRequest{}).DoAndReturn(func(req *TaskListRequest) error {
		req.Completed = &completed
		req.limit = 2
		req.offset = 1
		return nil
	})
	req, _ := NewTaskListRequest(ctx)
	assert.Equal(t, 2, req.GetLimit())
	assert.Equal(t, 1, req.GetOffset())
	assert.Equal(t, &completed, req.Completed)
}

func TestNewTaskCreateRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := mock_context.NewMockContext(ctrl)
	ctx.EXPECT().Bind(&TaskCreateRequest{}).DoAndReturn(func(req *TaskCreateRequest) error {
		req.Title = "foo"
		req.Description = "bar"
		return nil
	})
	req, _ := NewTaskCreateRequest(ctx)
	assert.Equal(t, "foo", req.Title)
	assert.Equal(t, "bar", req.Description)
}

func TestNewTaskUpdateRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := mock_context.NewMockContext(ctrl)
	ctx.EXPECT().Bind(&TaskUpdateRequest{}).DoAndReturn(func(req *TaskUpdateRequest) error {
		req.Title = "foo"
		req.Description = "bar"
		req.Completed = true
		return nil
	})
	ctx.EXPECT().Param("id").Return("1")
	req, _ := NewTaskUpdateRequest(ctx)
	assert.Equal(t, req.ID, 1)
	assert.Equal(t, "foo", req.Title)
	assert.Equal(t, "bar", req.Description)
	assert.Equal(t, true, req.Completed)
}
