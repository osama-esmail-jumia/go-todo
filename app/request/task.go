package request

import (
	"go-todo/context"
	"strconv"
)

type TaskListRequest struct {
	paginationRequest
	Completed *bool `form:"completed"`
}

type TaskCreateRequest struct {
	Title       string
	Description string
}

type TaskUpdateRequest struct {
	ID          int `json:"-"`
	Title       string
	Description string
	Completed   bool
}

func NewTaskListRequest(c context.Context) (req TaskListRequest, err error) {
	err = c.BindQuery(&req)
	return
}

func NewTaskCreateRequest(c context.Context) (req TaskCreateRequest, err error) {
	err = c.Bind(&req)
	return
}

func NewTaskUpdateRequest(c context.Context) (req TaskUpdateRequest, err error) {
	err = c.Bind(&req)
	req.ID, err = strconv.Atoi(c.Param("id"))
	return
}
