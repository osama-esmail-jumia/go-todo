package response

import "go-todo/app/model"

type TaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskListResponse struct {
	paginationResponse
	Rows []*TaskResponse `json:"rows"`
}

func NewTaskResponse(payload *model.Task) *TaskResponse {
	res := &TaskResponse{
		ID:          payload.ID,
		Title:       payload.Title,
		Description: payload.Description,
		Completed:   payload.Completed,
	}
	return res
}

func NewTaskListResponse(payloads []*model.Task, total, limit, offset int) *TaskListResponse {
	rows := make([]*TaskResponse, len(payloads))
	for i, payload := range payloads {
		rows[i] = NewTaskResponse(payload)
	}
	res := &TaskListResponse{
		paginationResponse: paginationResponse{
			Limit:  limit,
			Offset: offset,
			Total:  total,
		},
		Rows: rows,
	}

	return res
}
