// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Banner is the golang structure for table banner.
type Banner struct {
	Id          int         `json:"id"          ` // 横幅ID
	UserId      int         `json:"userId"      ` // 用户id
	Title       string      `json:"title"       ` // 横幅标题
	Info        string      `json:"info"        ` // 横幅信息
	CoverUrl    string      `json:"coverUrl"    ` // 横幅图片地址
	DisplayDate *gtime.Time `json:"displayDate" ` // 展示日期
	CreateAt    *gtime.Time `json:"createAt"    ` // 创建日期
	UpdateAt    *gtime.Time `json:"updateAt"    ` // 更新日期
}
