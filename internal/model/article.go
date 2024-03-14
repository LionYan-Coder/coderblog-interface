package model

import "coderblog-interface/internal/model/entity"

type ArticleCreateInput struct {
	Title   string
	Author  string
	Summary string
	Content string
}
type ArticleCreateOutput struct {
	Id int
}
type ArticleUpdateInput struct {
	Id      int
	Title   string
	Author  string
	Summary string
	Content string
}
type ArticleUpdateOutput struct {
}
type ArticleDeleteInput struct {
	Id int
}
type ArticleDeleteOutput struct {
}

type ArticleDetailInput struct {
	Id int
}
type ArticleDetailOutput struct {
	*entity.Article
}
type ArticleListInput struct {
	Page int
	Size int
}
type ArticleListOutput struct {
	List  []ArticleDetailOutput
	Total int `json:"total" dc:"总数"`
	Page  int `json:"page" dc:"分页号码"`
	Size  int `json:"size" dc:"分页数量"`
}

type ArticleListAllInput struct {
}
type ArticleListAllOutput struct {
	Total int
	List  []ArticleDetailOutput
}
