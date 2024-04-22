package adminV1

import (
	"coderblog-interface/api"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"
)

type ArticleBase struct {
	Title    string   `p:"title" v:"required|min-length:5|max-length:30#请填写标题|标题不能小于5个字符|标题不能超过30个字符" dc:"文章标题"`
	Summary  string   `p:"summary" v:"max-length:120#概要不能超过120个字符" dc:"文章概要"`
	Content  string   `p:"content" dc:"文章内容"`
	Tags     []string `p:"tags" dc:"文章标签"`
	CoverURL string   `p:"coverUrl" json:"coverUrl" v:"required|url#缺少文章封面|地址不合法" dc:"文章封面地址"`
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
	ID        int         `json:"id"        ` // 文章ID
	Author    string      `json:"author"    ` // 作者
	Title     string      `json:"title"     ` // 文章标题
	Summary   string      `json:"summary"   ` // 文章概要
	CoverURL  string      `json:"coverUrl"  ` // 文章图片
	Content   string      `json:"content"   ` // 文章内容
	Tags      []string    `json:"tags"      ` // 文章标签
	CreateAt  *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  ` // 更新时间
	Published bool        `json:"published" ` // 发布状态（1: 已发布，0: 未发布）
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
