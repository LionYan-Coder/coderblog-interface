// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id         int         `json:"id"         ` // 菜单id
	Name       string      `json:"name"       ` // 菜单名称
	Type       int         `json:"type"       ` // 菜单类型： （1: 模块，2: 页面，3: 页面按钮，4: 全局按钮  ）
	Comment    string      `json:"comment"    ` // 菜单描述
	Visibility int         `json:"visibility" ` // 菜单可见性 （0: 显示，1: 隐藏 ）
	Url        string      `json:"url"        ` // 页面路径： （菜单类型为页面时填写）
	ParentId   int         `json:"parentId"   ` // 父菜单id
	CreateAt   *gtime.Time `json:"createAt"   ` // 创建时间
	UpdateAt   *gtime.Time `json:"updateAt"   ` // 更新时间
}
