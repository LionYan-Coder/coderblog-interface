package v1

import (
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetRecentArticleByCurrentMonthReq struct {
	g.Meta `path:"/article/currentMonth" method:"get" tags:"内容服务" summary:"获取最近文章"`
}

type GetRecentArticleByCurrentMonthRes struct {
	List  []*entity.Article `json:"list" dc:"近期文章列表"`
	Total int               `json:"total" dc:"近期文章总数"`
}
