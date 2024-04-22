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
	g.Meta    `orm:"table:article, do:true"`
	Id        interface{} // 文章ID
	Author    interface{} // 作者
	AuthorId  interface{} // 作者id
	Title     interface{} // 文章标题
	Summary   interface{} // 文章概要
	CoverUrl  interface{} // 文章图片
	Content   interface{} // 文章内容
	Tags      interface{} // 文章标签
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	Published interface{} // 发布状态（1: 已发布，0: 未发布）
}
