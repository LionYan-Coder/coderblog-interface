package v1

import (
	"coderblog-interface/api"
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ArticleBase struct {
	Title   string `p:"title" v:"required|max-length:30#请填写标题|标题不能超过30个字符" dc:"文章标题"`
	Summary string `p:"summary" v:"max-length:120#概要不能超过120个字符" dc:"文章概要"`
	Content string `p:"content" dc:"文章内容"`
}

type CreateArticleReq struct {
	g.Meta `path:"/article" method:"put" tags:"内容服务" summary:"创建文章"`
	ArticleBase
}

type CreateArticleRes struct {
	ID int `json:"id" dc:"id"`
}

type UpdateArticleReq struct {
	g.Meta `path:"/article/{id}" method:"put,post" tags:"内容服务" summary:"修改文章"`
	ID     int `in:"path" v:"min:1#缺少文章ID" json:"id" dc:"id"`
	ArticleBase
}

type UpdateArticleRes struct {
}

type DeleteArticleReq struct {
	g.Meta `path:"/article/{id}" method:"delete" tags:"内容服务" summary:"删除文章"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type DeleteArticleRes struct {
}

type GetOneArticleReq struct {
	g.Meta `path:"/article/{id}" method:"get" tags:"内容服务" summary:"根据ID获取文章"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type GetOneArticleRes struct {
	*entity.Article
}

type GetListArticleReq struct {
	g.Meta `path:"/article" method:"get" tags:"内容服务" summary:"获取文章列表"`
	api.CommonPaginationReq
}

type GetListArticleRes struct {
	List  []GetOneArticleRes `json:"list" dc:"文章列表"`
	Total int                `json:"total" dc:"总数"`
	Page  int                `json:"page" dc:"分页号码"`
	Size  int                `json:"size" dc:"分页数量"`
}

type GetListAllArticleReq struct {
	g.Meta `path:"/article/all" method:"get" tags:"内容服务" summary:"获取所有文章"`
}

type GetListAllArticleRes struct {
	List  []GetOneArticleRes `json:"list" dc:"文章列表"`
	Total int                `json:"total" dc:"文章总数"`
}
