// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure of table article for DAO operations like Where/Data.
type Article struct {
	g.Meta   `orm:"table:article, do:true"`
	Id       interface{} // 文章ID
	Author   interface{} // 作者
	Title    interface{} // 文章标题
	Summary  interface{} // 文章概要
	Content  interface{} // 文章内容
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
