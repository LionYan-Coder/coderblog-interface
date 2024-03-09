// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta   `orm:"table:role, do:true"`
	Id       interface{} // 角色ID
	Name     interface{} // 角色名称
	Type     interface{} // 角色类型： 0(管理员) 1（用户）
	Comment  interface{} // 角色注释
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
