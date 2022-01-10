package request

import "go-todo/config"

type paginationRequest struct {
	limit  int `form:"limit"`
	offset int `form:"offset"`
}

func (req *paginationRequest) GetLimit() int {
	if req.limit == 0 {
		return config.Cfg().AppDefaultPageSize
	}

	return req.limit
}

func (req *paginationRequest) GetOffset() int {
	return req.offset
}
