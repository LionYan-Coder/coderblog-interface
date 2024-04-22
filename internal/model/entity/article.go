// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure for table article.
type Article struct {
	Id        int         `json:"id"        ` // 文章ID
	Author    string      `json:"author"    ` // 作者
	AuthorId  string      `json:"authorId"  ` // 作者id
	Title     string      `json:"title"     ` // 文章标题
	Summary   string      `json:"summary"   ` // 文章概要
	CoverUrl  string      `json:"coverUrl"  ` // 文章图片
	Content   string      `json:"content"   ` // 文章内容
	Tags      string      `json:"tags"      ` // 文章标签
	CreateAt  *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  ` // 更新时间
	Published int         `json:"published" ` // 发布状态（1: 已发布，0: 未发布）
}
