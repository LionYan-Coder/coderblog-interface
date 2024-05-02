// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleComment is the golang structure of table article_comment for DAO operations like Where/Data.
type ArticleComment struct {
	g.Meta    `orm:"table:article_comment, do:true"`
	Id        interface{} // 评论id
	UserId    interface{} // 用户id
	ArticleId interface{} // 评论的文章id
	ReplyId   interface{} // 回复评论id
	Content   interface{} // 评论内容
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
}
