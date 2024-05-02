// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NotebookComment is the golang structure of table notebook_comment for DAO operations like Where/Data.
type NotebookComment struct {
	g.Meta     `orm:"table:notebook_comment, do:true"`
	Id         interface{} // 笔记id
	UserId     interface{} // 用户id
	NotebookId interface{} // 笔记id
	ReplyId    interface{} // 回复id
	Content    interface{} // 内容
	CreateAt   *gtime.Time // 创建时间
	UpdateAt   *gtime.Time // 更新时间
}
