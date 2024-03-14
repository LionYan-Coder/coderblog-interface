package v1

import (
	"coderblog-interface/api"
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PermissionBase struct {
	Key     string `p:"key" v:"required|max-length:30#缺少权限Key|权限Key不能超过30个字符" dc:"权限Key"`
	Comment string `p:"comment" v:"max-length:120#注释不能超过120个字符" dc:"权限注释"`
}

type CreatePermissionReq struct {
	g.Meta `path:"/permission" method:"put" tags:"权限服务" summary:"创建权限"`
	PermissionBase
}

type CreatePermissionRes struct {
	ID int `json:"id" dc:"id"`
}

type UpdatePermissionReq struct {
	g.Meta `path:"/permission/{id}" method:"put,post" tags:"权限服务" summary:"修改权限"`
	ID     int `in:"path" v:"min:1#缺少权限ID" json:"id" dc:"id"`
	PermissionBase
}

type UpdatePermissionRes struct {
}

type DeletePermissionReq struct {
	g.Meta `path:"/permission/{id}" method:"delete" tags:"权限服务" summary:"删除权限"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type DeletePermissionRes struct {
}

type GetOnePermissionReq struct {
	g.Meta `path:"/permission/{id}" method:"get" tags:"权限服务" summary:"获取权限"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type GetOnePermissionRes struct {
	*entity.Permission
}

type GetListPermissionReq struct {
	g.Meta `path:"/permission" method:"get" tags:"权限服务" summary:"获取权限"`
	api.CommonPaginationReq
}

type GetListPermissionRes struct {
	List  []GetOnePermissionRes `json:"list" dc:"权限列表"`
	Total int                   `json:"total" dc:"总数"`
	Page  int                   `json:"page" dc:"分页号码"`
	Size  int                   `json:"size" dc:"分页数量"`
}

type GetListAllPermissionReq struct {
	g.Meta `path:"/permission/all" method:"get" tags:"权限服务" summary:"获取所有权限"`
}

type GetListAllPermissionRes struct {
	List  []GetOnePermissionRes `json:"list" dc:"权限列表"`
	Total int                   `json:"total" dc:"权限总数"`
}
