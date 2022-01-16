package request

import (
	"go-todo/context"
	"go-todo/pkg/validation"
	"strconv"
)

type TaskListRequest struct {
	paginationRequest
	Completed *bool `form:"completed"`
}

type TaskCreateRequest struct {
	Title       string `validate:"required,min=2,max=100"`
	Description string `validate:"required,min=4,max=600"`
}

type TaskUpdateRequest struct {
	ID          int    `json:"-"`
	Title       string `validate:"omitempty,min=2,max=100"`
	Description string `validate:"omitempty,min=4,max=600"`
	Completed   bool   `validate:"required"`
}

func NewTaskListRequest(c context.Context) (req TaskListRequest, err error) {
	err = c.BindQuery(&req)
	return
}

func NewTaskCreateRequest(c context.Context) (req TaskCreateRequest, err error) {
	err = c.Bind(&req)
	if err != nil {
		return
	}

	err = validation.Validate().Struct(req)
	return
}

func NewTaskUpdateRequest(c context.Context) (req TaskUpdateRequest, err error) {
	err = c.Bind(&req)
	if err != nil {
		return
	}

	req.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	err = validation.Validate().Struct(req)
	return
}
