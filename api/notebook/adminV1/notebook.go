package adminV1

import (
	"coderblog-interface/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type NotebookDTO struct {
	Title   string   `p:"title" json:"title" v:"required|min-length:5|max-length:30#请填写标题|标题不能小于5个字符|标题不能超过30个字符" dc:"笔记标题"`
	Summary string   `p:"summary" json:"summary" v:"max-length:120#概要不能超过120个字符" dc:"笔记概要"`
	Content string   `p:"content" json:"content" dc:"笔记内容"`
	Tags    []string `p:"tags" json:"tags" dc:"笔记标签"`
}

type CreateNotebookReq struct {
	g.Meta `path:"/notebook" method:"put" tags:"笔记服务" summary:"创建笔记"`
	NotebookDTO
}

type CreateNotebookRes struct {
	ID int `json:"id" dc:"id"`
}

type UpdateNotebookReq struct {
	g.Meta `path:"/notebook/{id}" method:"post" tags:"笔记服务" summary:"修改笔记"`
	ID     int `in:"path" v:"min:1#缺少笔记ID" json:"id" dc:"id"`
	NotebookDTO
}

type UpdateNotebookRes struct {
}

type DeleteNotebookReq struct {
	g.Meta `path:"/notebook/{id}" method:"delete" tags:"笔记服务" summary:"删除笔记"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type DeleteNotebookRes struct {
}

type GetOneNotebookReq struct {
	g.Meta `path:"/notebook/{id}" method:"get" tags:"笔记服务" summary:"根据ID获取笔记"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type GetOneNotebookRes struct {
	NotebookDTO
	ID        int         `json:"id"        ` // 笔记ID
	Author    string      `json:"author"    ` // 作者
	UserID    string      `json:"userId"    ` // 作者id
	CreateAt  *gtime.Time `json:"createAt"  ` // 创建时间
	UpdateAt  *gtime.Time `json:"updateAt"  ` // 更新时间
	Published bool        `json:"published" ` // 发布状态（1: 已发布，0: 未发布）
}

type GetListNotebookReq struct {
	g.Meta `path:"/notebook" method:"get" tags:"笔记服务" summary:"获取笔记列表"`
	api.CommonPaginationReq
}

type GetListNotebookRes struct {
	List  []GetOneNotebookRes `json:"list" dc:"笔记列表"`
	Total int                 `json:"total" dc:"总数"`
	Page  int                 `json:"page" dc:"分页号码"`
	Size  int                 `json:"size" dc:"分页数量"`
}

type PublishNotebookReq struct {
	g.Meta `path:"/notebook/publish/{id}" method:"post" tags:"笔记服务" summary:"发布笔记"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type PublishNotebookRes struct {
}

type UnPublishNotebookReq struct {
	g.Meta `path:"/notebook/unpublish/{id}" method:"post" tags:"笔记服务" summary:"撤销发布笔记"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type UnPublishNotebookRes struct {
}
