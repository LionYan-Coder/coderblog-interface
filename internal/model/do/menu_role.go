// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuRole is the golang structure of table menu_role for DAO operations like Where/Data.
type MenuRole struct {
	g.Meta   `orm:"table:menu_role, do:true"`
	Id       interface{} // 菜单-角色关联id
	MenuId   interface{} // 菜单id
	RoleId   interface{} // 角色id
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
