// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notebook is the golang structure of table notebook for DAO operations like Where/Data.
type Notebook struct {
	g.Meta    `orm:"table:notebook, do:true"`
	Id        interface{} //
	UserId    interface{} // 用户id
	Title     interface{} // 笔记标题
	Summary   interface{} // 笔记概要
	Tags      interface{} // 笔记标签
	Content   interface{} // 笔记内容
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	Published interface{} // 发布状态（1: 已发布，0: 未发布）
}
