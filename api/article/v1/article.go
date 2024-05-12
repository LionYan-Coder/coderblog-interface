package v1

import (
	"coderblog-interface/api"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"
)

type ArticleDTO struct {
	ID       int         `json:"id"        ` // 文章ID
	Author   string      `json:"author"    ` // 作者
	Title    string      `json:"title"     ` // 文章标题
	Summary  string      `json:"summary"   ` // 文章概要
	CoverURL string      `json:"coverUrl"  ` // 文章图片
	Content  string      `json:"content"   ` // 文章内容
	Tags     []string    `json:"tags"      ` // 文章标签
	CreateAt *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt"  ` // 更新时间
}

type GetListArticleReq struct {
	g.Meta `path:"/article" method:"get" tags:"前端-博客服务" summary:"获取文章列表"`
	api.CommonPaginationReq
}

type GetListArticleRes struct {
	List  []ArticleDTO `json:"list" dc:"文章列表"`
	Total int          `json:"total" dc:"总数"`
	Page  int          `json:"page" dc:"分页号码"`
	Size  int          `json:"size" dc:"分页数量"`
}

type GetOneArticleReq struct {
	g.Meta `path:"/article/{title}" method:"get" tags:"前端-博客服务" summary:"根据标题获取文章"`
	Title  string `in:"path" json:"title" dc:"title"`
}

type GetOneArticleRes struct {
	ArticleDTO
}
