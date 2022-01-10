package request

import "go-todo/config"

type paginationRequest struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

func (req *paginationRequest) GetLimit() int {
	if req.Limit == 0 {
		return config.Cfg().AppDefaultPageSize
	}

	return req.Limit
}

func (req *paginationRequest) GetOffset() int {
	return req.Offset
}
