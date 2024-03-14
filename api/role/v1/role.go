package v1

import (
	"coderblog-interface/api"
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateRoleReq struct {
	g.Meta  `path:"/role" method:"put" tags:"角色服务" summary:"创建角色"`
	Name    string `p:"name" v:"required|min-length:2|max-length:16#请填写角色名称|角色名称至少2个字符|角色名称不能超过16个字符"`
	Type    int    `p:"type" v:"in:0,1#角色类型错误" dc:"角色类型： 0(管理员) 1（用户）" d:"1"`
	Comment string `p:"comment" dc:"角色注释"`
}

type CreateRoleRes struct {
	ID int `json:"id" dc:"id"`
}

type UpdateRoleReq struct {
	g.Meta  `path:"/role/{id}" method:"put,post" tags:"角色服务" summary:"修改角色"`
	ID      int    `in:"path" v:"min:1#缺少角色Id" dc:"角色ID"`
	Name    string `p:"name" v:"min-length:2|max-length:16#角色名称至少2个字符|角色名称不能超过16个字符"`
	Type    int    `p:"type" v:"in:0,1#角色类型不合法" dc:"角色类型： 0(管理员) 1（用户）" d:"1"`
	Comment string `p:"comment" dc:"角色注释"`
}

type UpdateRoleRes struct {
}

type DeleteRoleReq struct {
	g.Meta `path:"/role/{id}" method:"delete" tags:"角色服务" summary:"删除角色"`
	ID     int `in:"path" v:"min:1#缺少角色Id" dc:"角色ID"`
}

type DeleteRoleRes struct {
}

type GetOneRoleReq struct {
	g.Meta `path:"/role/{id}" method:"get" tags:"角色服务" summary:"获取角色"`
	ID     int `in:"path" v:"min:1#缺少角色Id" dc:"角色ID"`
}

type GetOneRoleRes struct {
	*entity.Role
}

type GetListRoleReq struct {
	g.Meta `path:"/role" method:"get" tags:"角色服务" summary:"通过分页获取角色"`
	api.CommonPaginationReq
}

type GetListRoleRes struct {
	List  []GetOneRoleRes `json:"list" dc:"角色列表数据"`
	Total int             `json:"total" dc:"总数"`
	Page  int             `json:"page" dc:"分页号码"`
	Size  int             `json:"size" dc:"分页数量"`
}

type GetAllRoleReq struct {
	g.Meta `path:"/role/all" method:"get" tags:"角色服务" summary:"获取全部角色"`
}

type GetAllRoleRes struct {
	List  []GetOneRoleRes `json:"list" dc:"角色列表数据"`
	Total int             `json:"total" dc:"总数"`
}
