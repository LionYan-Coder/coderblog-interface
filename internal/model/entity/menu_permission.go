// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuPermission is the golang structure for table menu_permission.
type MenuPermission struct {
	Id           int         `json:"id"           ` // 菜单-权限 关联id
	MenuId       int         `json:"menuId"       ` // 菜单id
	PermissionId int         `json:"permissionId" ` // 权限id
	CreateAt     *gtime.Time `json:"createAt"     ` // 创建时间
	UpdateAt     *gtime.Time `json:"updateAt"     ` // 更新时间
}
