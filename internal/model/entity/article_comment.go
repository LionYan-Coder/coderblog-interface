// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleComment is the golang structure for table article_comment.
type ArticleComment struct {
	Id        int         `json:"id"        ` // 评论id
	UserId    string      `json:"userId"    ` // 用户id
	ArticleId int         `json:"articleId" ` // 评论的文章id
	ReplyId   int         `json:"replyId"   ` // 回复评论id
	Content   string      `json:"content"   ` // 评论内容
	CreateAt  *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  ` // 更新时间
}
