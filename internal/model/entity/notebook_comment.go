// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NotebookComment is the golang structure for table notebook_comment.
type NotebookComment struct {
	Id         int         `json:"id"         ` // 笔记id
	UserId     string      `json:"userId"     ` // 用户id
	NotebookId int         `json:"notebookId" ` // 笔记id
	ReplyId    int         `json:"replyId"    ` // 回复id
	Content    string      `json:"content"    ` // 内容
	CreateAt   *gtime.Time `json:"createAt"   ` // 创建时间
	UpdateAt   *gtime.Time `json:"updateAt"   ` // 更新时间
}
