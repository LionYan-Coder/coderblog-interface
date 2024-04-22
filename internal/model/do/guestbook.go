// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Guestbook is the golang structure of table guestbook for DAO operations like Where/Data.
type Guestbook struct {
	g.Meta   `orm:"table:guestbook, do:true"`
	Id       interface{} //
	Author   interface{} // 留言人
	Content  interface{} // 留言内容
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
