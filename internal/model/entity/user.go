// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       int         `json:"id"       ` // 用户id
	Username string      `json:"username" ` // 用户名
	Password string      `json:"password" ` // 用户密码
	Nickname string      `json:"nickname" ` // 用户昵称
	CreateAt *gtime.Time `json:"createAt" ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新时间
}
