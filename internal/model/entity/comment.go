// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comment is the golang structure for table comment.
type Comment struct {
	Id        int         `json:"id"        ` // 评论id
	AuthorId  string      `json:"authorId"  ` // 用户id
	ArticleId int         `json:"articleId" ` // 评论的文章id
	ReplyId   int         `json:"replyId"   ` // 回复评论id
	Content   string      `json:"content"   ` // 评论内容
	CreateAt  *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  ` // 更新时间
}
