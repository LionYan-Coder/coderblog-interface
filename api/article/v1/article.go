package v1

import (
	"coderblog-interface/api"
	"coderblog-interface/api/article/adminV1"

	"github.com/gogf/gf/v2/frame/g"
)

type GetRecentArticleByCurrentMonthReq struct {
	g.Meta `path:"/article/currentMonth" method:"get" tags:"内容服务" summary:"获取最近文章"`
}

type GetRecentArticleByCurrentMonthRes struct {
	List  []*adminV1.GetOneArticleRes `json:"list" dc:"近期文章列表"`
	Total int                         `json:"total" dc:"近期文章总数"`
}

type GetListArticleReq struct {
	g.Meta `path:"/article" method:"get" tags:"内容服务" summary:"获取文章列表"`
	api.CommonPaginationReq
}

type GetListArticleRes struct {
	List  []*adminV1.GetOneArticleRes `json:"list" dc:"文章列表"`
	Total int                         `json:"total" dc:"总数"`
	Page  int                         `json:"page" dc:"分页号码"`
	Size  int                         `json:"size" dc:"分页数量"`
}
