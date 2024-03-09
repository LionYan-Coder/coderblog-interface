// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id       int         `json:"id"       ` // 角色ID
	Name     string      `json:"name"     ` // 角色名称
	Type     int         `json:"type"     ` // 角色类型： 0(管理员) 1（用户）
	Comment  string      `json:"comment"  ` // 角色注释
	CreateAt *gtime.Time `json:"createAt" ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新时间
}
