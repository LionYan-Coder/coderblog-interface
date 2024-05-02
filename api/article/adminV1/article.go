package adminV1

import (
	"coderblog-interface/api"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"
)

type ArticleDTO struct {
	Title    string   `p:"title" json:"title" v:"required|min-length:5|max-length:30#请填写标题|标题不能小于5个字符|标题不能超过30个字符" dc:"文章标题"`
	Summary  string   `p:"summary" json:"summary" v:"max-length:120#概要不能超过120个字符" dc:"文章概要"`
	Content  string   `p:"content" json:"content" dc:"文章内容"`
	Tags     []string `p:"tags" json:"tags" dc:"文章标签"`
	CoverURL string   `p:"coverUrl" json:"coverUrl" v:"required|url#缺少文章封面|地址不合法" dc:"文章封面地址"`
}

type CreateArticleReq struct {
	g.Meta `path:"/article" method:"put" tags:"博客文章" summary:"创建文章"`
	ArticleDTO
}

type CreateArticleRes struct {
	ID int `json:"id" dc:"id"`
}

type UpdateArticleReq struct {
	g.Meta `path:"/article/{id}" method:"post" tags:"博客文章" summary:"修改文章"`
	ID     int `in:"path" v:"min:1#缺少文章ID" json:"id" dc:"id"`
	ArticleDTO
}

type UpdateArticleRes struct {
}

type DeleteArticleReq struct {
	g.Meta `path:"/article/{id}" method:"delete" tags:"博客文章" summary:"删除文章"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type DeleteArticleRes struct {
}

type GetOneArticleReq struct {
	g.Meta `path:"/article/{id}" method:"get" tags:"博客文章" summary:"根据ID获取文章"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type GetOneArticleRes struct {
	ArticleDTO
	ID        int         `json:"id"        ` // 文章ID
	Author    string      `json:"author"    ` // 作者
	UserID    string      `json:"userId"    ` // 作者id
	Tags      []string    `json:"tags"      ` // 文章标签
	CreateAt  *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  ` // 更新时间
	Published bool        `json:"published" ` // 发布状态（1: 已发布，0: 未发布）
}

type GetListArticleReq struct {
	g.Meta `path:"/article" method:"get" tags:"博客文章" summary:"获取文章列表"`
	api.CommonPaginationReq
}

type GetListArticleRes struct {
	List  []GetOneArticleRes `json:"list" dc:"文章列表"`
	Total int                `json:"total" dc:"总数"`
	Page  int                `json:"page" dc:"分页号码"`
	Size  int                `json:"size" dc:"分页数量"`
}

type PublishArticleReq struct {
	g.Meta `path:"/article/publish/{id}" method:"post" tags:"博客文章" summary:"发布文章"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type PublishArticleRes struct {
}

type UnPublishArticleReq struct {
	g.Meta `path:"/article/unpublish/{id}" method:"post" tags:"博客文章" summary:"撤销发布文章"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type UnPublishArticleRes struct {
}
