package v1

import (
	"coderblog-interface/api"
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type MenuBase struct {
	Name       string `p:"name" v:"required|max-length:30#请填写名称|名称不能超过30个字符" dc:"菜单名称"`
	Type       int    `p:"type" v:"in:1,2,3,4#菜单类型错误" d:"1" dc:"菜单类型:（1: 模块，2: 页面，3: 页面按钮，4: 全局按钮 ）"`
	Comment    string `p:"comment" v:"max-length:120#注释不能超过120个字符" dc:"菜单注释"`
	Visibility int    `p:"visibility" v:"in:0,1#可见性类型错误" d:"0" dc:"菜单可见性:（0: 显示，1: 隐藏 ）"`
	Url        string `p:"url" v:"max-length:30#url不能超过30个字符" dc:"页面路径:（菜单类型为页面时填写）"`
	ParentId   int    `p:"parentId" dc:"父菜单ID (顶层菜单不传为空)"`
}

type CreateMenuReq struct {
	g.Meta `path:"/menu" method:"put" tags:"权限服务" summary:"创建菜单"`
	MenuBase
}

type CreateMenuRes struct {
	ID int `json:"id" dc:"id"`
}

type UpdateMenuReq struct {
	g.Meta `path:"/menu/{id}" method:"put,post" tags:"权限服务" summary:"修改菜单"`
	ID     int `in:"path" v:"min:1#缺少菜单ID" json:"id" dc:"id"`
	MenuBase
}

type UpdateMenuRes struct {
}

type DeleteMenuReq struct {
	g.Meta `path:"/menu/{id}" method:"delete" tags:"权限服务" summary:"删除菜单"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type DeleteMenuRes struct {
}

type GetOneMenuReq struct {
	g.Meta `path:"/menu/{id}" method:"get" tags:"权限服务" summary:"获取菜单"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type GetOneMenuRes struct {
	*entity.Menu
}

type GetListMenuReq struct {
	g.Meta `path:"/menu" method:"get" tags:"权限服务" summary:"获取菜单"`
	api.CommonPaginationReq
}

type GetListMenuRes struct {
	List  []GetOneMenuRes `json:"list" dc:"菜单列表"`
	Total int             `json:"total" dc:"总数"`
	Page  int             `json:"page" dc:"分页号码"`
	Size  int             `json:"size" dc:"分页数量"`
}

type GetListAllMenuReq struct {
	g.Meta `path:"/menu/all" method:"get" tags:"权限服务" summary:"获取所有菜单"`
}

type GetListAllMenuRes struct {
	List  []GetOneMenuRes `json:"list" dc:"菜单列表"`
	Total int             `json:"total" dc:"菜单总数"`
}
