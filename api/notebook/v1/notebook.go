package v1

import (
	"coderblog-interface/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type NotebookDTO struct {
	ID       int         `json:"id"        ` // 笔记ID
	Author   string      `json:"author"    ` // 作者
	Title    string      `json:"title"     ` // 笔记标题
	Summary  string      `json:"summary"   ` // 笔记概要
	Content  string      `json:"content"   ` // 笔记内容
	Tags     []string    `json:"tags"      ` // 笔记标签
	CreateAt *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt"  ` // 更新时间
}

type GetListNotebookReq struct {
	g.Meta `path:"/notebook" method:"get" tags:"前端-笔记服务" summary:"获取笔记列表"`
	api.CommonPaginationReq
}

type GetListNotebookRes struct {
	List  []GetOneNotebookRes `json:"list" dc:"笔记列表"`
	Total int                 `json:"total" dc:"总数"`
	Page  int                 `json:"page" dc:"分页号码"`
	Size  int                 `json:"size" dc:"分页数量"`
}

type GetOneNotebookReq struct {
	g.Meta `path:"/notebook/{id}" method:"get" tags:"前端-笔记服务" summary:"根据id获取笔记"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type GetOneNotebookRes struct {
	NotebookDTO
}
