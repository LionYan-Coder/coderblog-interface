// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Guestbook is the golang structure for table guestbook.
type Guestbook struct {
	Id       int         `json:"id"       ` //
	Author   string      `json:"author"   ` // 留言人
	Content  string      `json:"content"  ` // 留言内容
	CreateAt *gtime.Time `json:"createAt" ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新时间
}
