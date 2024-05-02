package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type NotebookCreateInput struct {
	UserID  string   `json:"userId" `
	Title   string   `json:"title"     ` // 笔记标题
	Summary string   `json:"summary"   ` // 笔记概要
	Content string   `json:"content"   ` // 笔记内容
	Tags    []string `json:"tags"      ` // 笔记标签
}
type NotebookCreateOutput struct {
	ID int
}
type NotebookUpdateInput struct {
	UserID  string   `json:"userId" `
	Title   string   `json:"title"     ` // 笔记标题
	Summary string   `json:"summary"   ` // 笔记概要
	Content string   `json:"content"   ` // 笔记内容
	Tags    []string `json:"tags"      ` // 笔记标签
	ID      int      `json:"id"        ` // 笔记ID
}
type NotebookUpdateOutput struct {
}
type NotebookDeleteInput struct {
	ID int
}
type NotebookDeleteOutput struct {
}

type NotebookDetailInput struct {
	ID int
}
type NotebookDetailOutput struct {
	NotebookUpdateInput
	CreateAt  *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  ` // 更新时间
	Published bool        `json:"published" ` // 发布状态（1: 已发布，0: 未发布）
}
type NotebookListInput struct {
	Page int
	Size int
}
type NotebookListOutput struct {
	List  []NotebookDetailOutput
	Total int `json:"total" dc:"总数"`
	Page  int `json:"page" dc:"分页号码"`
	Size  int `json:"size" dc:"分页数量"`
}

type NotebookPublishInput struct {
	ID int
}

type NotebookPublishOutput struct{}

type NotebookUnPublishInput struct {
	ID int
}

type NotebookUnPublishOutput struct{}
