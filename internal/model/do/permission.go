// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Permission is the golang structure of table permission for DAO operations like Where/Data.
type Permission struct {
	g.Meta   `orm:"table:permission, do:true"`
	Id       interface{} // 权限id
	Key      interface{} // 权限key
	Comment  interface{} // 权限描述
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
