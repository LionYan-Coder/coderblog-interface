// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRole is the golang structure for table user_role.
type UserRole struct {
	Id       int         `json:"id"       ` // ID
	UserId   int         `json:"userId"   ` // 用户id
	RoleId   int         `json:"roleId"   ` // 角色id
	CreateAt *gtime.Time `json:"createAt" ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新时间
}
