// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta     `orm:"table:menu, do:true"`
	Id         interface{} // 菜单id
	Name       interface{} // 菜单名称
	Type       interface{} // 菜单类型： （1: 模块，2: 页面，3: 页面按钮，4: 全局按钮  ）
	Comment    interface{} // 菜单描述
	Visibility interface{} // 菜单可见性 （0: 显示，1: 隐藏 ）
	Url        interface{} // 页面路径： （菜单类型为页面时填写）
	ParentId   interface{} // 父菜单id
	CreateAt   *gtime.Time // 创建时间
	UpdateAt   *gtime.Time // 更新时间
}
