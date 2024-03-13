package model

type CommonPaginationInput struct {
	Page int
	Size int
}

type CommonPaginationOutput struct {
	Total int `json:"total" dc:"总数"`
	Page  int `json:"page" dc:"分页号码"`
	Size  int `json:"size" dc:"分页数量"`
}
