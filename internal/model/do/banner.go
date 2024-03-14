// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Banner is the golang structure of table banner for DAO operations like Where/Data.
type Banner struct {
	g.Meta      `orm:"table:banner, do:true"`
	Id          interface{} // 横幅ID
	UserId      interface{} // 用户id
	Title       interface{} // 横幅标题
	Info        interface{} // 横幅信息
	CoverUrl    interface{} // 横幅图片地址
	DisplayDate *gtime.Time // 展示日期
	CreateAt    *gtime.Time // 创建日期
	UpdateAt    *gtime.Time // 更新日期
}
