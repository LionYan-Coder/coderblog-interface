// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:user, do:true"`
	Id       interface{} // 用户id
	Username interface{} // 用户名
	Password interface{} // 用户密码
	Nickname interface{} // 用户昵称
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
