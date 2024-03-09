// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRole is the golang structure of table user_role for DAO operations like Where/Data.
type UserRole struct {
	g.Meta   `orm:"table:user_role, do:true"`
	Id       interface{} // ID
	UserId   interface{} // 用户id
	RoleId   interface{} // 角色id
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
