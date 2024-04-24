package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type ArticleCreateInput struct {
	Title    string
	Author   string
	AuthorID string
	Summary  string
	Content  string
	CoverURL string
	Tags     []string
}
type ArticleCreateOutput struct {
	ID int
}
type ArticleUpdateInput struct {
	ID       int
	Title    string
	Author   string
	AuthorID string
	Summary  string
	Content  string
	CoverURL string
	Tags     []string
}
type ArticleUpdateOutput struct {
}
type ArticleDeleteInput struct {
	ID int
}
type ArticleDeleteOutput struct {
}

type ArticleDetailInput struct {
	ID int
}
type ArticleDetailOutput struct {
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

type ArticlePublishInput struct {
	ID int
}

type ArticlePublishOutput struct{}

type ArticleUnPublishInput struct {
	ID int
}

type ArticleUnPublishOutput struct{}

type ArticleListAllInput struct {
}
type ArticleListAllOutput struct {
	Total int
	List  []ArticleDetailOutput
}

type ArticleGetRecentByCurrentMonthInput struct {
}
