// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuPermission is the golang structure of table menu_permission for DAO operations like Where/Data.
type MenuPermission struct {
	g.Meta       `orm:"table:menu_permission, do:true"`
	Id           interface{} // 菜单-权限 关联id
	MenuId       interface{} // 菜单id
	PermissionId interface{} // 权限id
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 更新时间
}
