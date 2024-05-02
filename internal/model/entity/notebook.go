// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notebook is the golang structure for table notebook.
type Notebook struct {
	Id        int         `json:"id"        ` //
	UserId    string      `json:"userId"    ` // 用户id
	Title     string      `json:"title"     ` // 笔记标题
	Summary   string      `json:"summary"   ` // 笔记概要
	Tags      string      `json:"tags"      ` // 笔记标签
	Content   string      `json:"content"   ` // 笔记内容
	CreateAt  *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  ` // 更新时间
	Published int         `json:"published" ` // 发布状态（1: 已发布，0: 未发布）
}
