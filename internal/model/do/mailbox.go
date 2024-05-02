// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Mailbox is the golang structure of table mailbox for DAO operations like Where/Data.
type Mailbox struct {
	g.Meta   `orm:"table:mailbox, do:true"`
	Id       interface{} //
	UserId   interface{} // 用户id
	Content  interface{} // 留言内容
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
