// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Mailbox is the golang structure for table mailbox.
type Mailbox struct {
	Id       int         `json:"id"       ` //
	UserId   string      `json:"userId"   ` // 用户id
	Content  string      `json:"content"  ` // 留言内容
	CreateAt *gtime.Time `json:"createAt" ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新时间
}
