// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuRole is the golang structure for table menu_role.
type MenuRole struct {
	Id       int         `json:"id"       ` // 菜单-角色关联id
	MenuId   int         `json:"menuId"   ` // 菜单id
	RoleId   int         `json:"roleId"   ` // 角色id
	CreateAt *gtime.Time `json:"createAt" ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新时间
}
