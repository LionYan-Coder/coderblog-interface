// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Note is the golang structure of table note for DAO operations like Where/Data.
type Note struct {
	g.Meta   `orm:"table:note, do:true"`
	Id       interface{} //
	Author   interface{} // 笔记作者
	Title    interface{} // 笔记标题
	Summary  interface{} // 笔记概要
	Content  interface{} // 笔记内容
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
