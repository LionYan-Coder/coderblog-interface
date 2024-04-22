// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Note is the golang structure for table note.
type Note struct {
	Id       int         `json:"id"       ` //
	Author   string      `json:"author"   ` // 笔记作者
	Title    string      `json:"title"    ` // 笔记标题
	Summary  string      `json:"summary"  ` // 笔记概要
	Content  string      `json:"content"  ` // 笔记内容
	CreateAt *gtime.Time `json:"createAt" ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新时间
}
