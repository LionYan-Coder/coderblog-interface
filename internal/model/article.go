package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type ArticleCreateInput struct {
	UserID   string   `json:"userId" `
	Title    string   `json:"title"     ` // 文章标题
	Summary  string   `json:"summary"   ` // 文章概要
	CoverURL string   `json:"coverUrl"  ` // 文章图片
	Content  string   `json:"content"   ` // 文章内容
	Tags     []string `json:"tags"      ` // 文章标签
}
type ArticleCreateOutput struct {
	ID int
}
type ArticleUpdateInput struct {
	UserID   string   `json:"userId" `
	Title    string   `json:"title"     ` // 文章标题
	Summary  string   `json:"summary"   ` // 文章概要
	CoverURL string   `json:"coverUrl"  ` // 文章图片
	Content  string   `json:"content"   ` // 文章内容
	Tags     []string `json:"tags"      ` // 文章标签
	ID       int      `json:"id"        ` // 文章ID
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
	ArticleUpdateInput
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
