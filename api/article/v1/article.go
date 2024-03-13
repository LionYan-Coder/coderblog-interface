package v1

import (
	"coderblog-interface/api"
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ArticleBase struct {
	Title   string `v:"required|max-length:30#请填写标题|标题不能超过30个字符" dc:"文章标题"`
	Summary string `v:"max-length:120#概要不能超过120个字符" dc:"文章概要"`
	Content string `dc:"文章内容"`
}

type CreateArticleReq struct {
	g.Meta `path:"/article" method:"put" tags:"内容服务" summary:"创建文章"`
	ArticleBase
}

type CreateArticleRes struct {
	Id int `json:"id" dc:"id"`
}

type UpdateArticleReq struct {
	g.Meta `path:"/article/{Id}" method:"put,post" tags:"内容服务" summary:"修改文章"`
	Id     int `in:"path" json:"id" dc:"id"`
	ArticleBase
}

type UpdateArticleRes struct {
}

type DeleteArticleReq struct {
	g.Meta `path:"/article/{Id}" method:"delete" tags:"内容服务" summary:"删除文章"`
	Id     int `in:"path" json:"id" dc:"id"`
}

type DeleteArticleRes struct {
}

type GetOneArticleReq struct {
	g.Meta `path:"/article/{Id}" method:"get" tags:"内容服务" summary:"获取文章"`
	Id     int `in:"path" json:"id" dc:"id"`
}

type GetOneArticleRes struct {
	*entity.Article
}

type GetListArticleReq struct {
	g.Meta `path:"/article" method:"get" tags:"内容服务" summary:"获取文章"`
	api.CommonPaginationReq
}

type GetListArticleRes struct {
	List []GetOneArticleRes `json:"list" dc:"文章列表"`
	api.CommonPaginationRes
}

type GetListAllArticleReq struct {
	g.Meta `path:"/article/all" method:"get" tags:"内容服务" summary:"获取所有文章"`
}

type GetListAllArticleRes struct {
	List  []GetOneArticleRes `json:"list" dc:"文章列表"`
	Total int                `json:"total" dc:"文章总数"`
}
