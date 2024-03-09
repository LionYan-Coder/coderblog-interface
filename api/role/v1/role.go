package v1

import (
	"coderblog-interface/api"
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateRoleReq struct {
	g.Meta  `path:"/role" method:"put" tags:"角色服务" summary:"创建角色"`
	Name    string `p:"name" v:"required|min:2|max:16#请填写角色名称|角色名称至少2个字符|角色名称不能超过16个字符"`
	Type    uint   `p:"type" v:"in:0,1#角色类型不合法" dc:"角色类型： 0(管理员) 1（用户）"`
	Comment string `p:"comment" dc:"角色注释"`
}

type CreateRoleRes struct {
	ID uint `json:"id" dc:"id"`
}

type UpdateRoleReq struct {
	g.Meta  `path:"/role/{id}" method:"put,post" tags:"角色服务" summary:"修改角色"`
	ID      uint   `p:"id" v:"required#缺少角色ID" in:"path"`
	Name    string `p:"name" v:"min:2|max:16#角色名称至少2个字符|角色名称不能超过16个字符"`
	Type    uint   `p:"type" v:"in:0,1#角色类型不合法" dc:"角色类型： 0(管理员) 1（用户）"`
	Comment string `p:"comment" dc:"角色注释"`
}

type UpdateRoleRes struct {
}

type DeleteRoleReq struct {
	g.Meta `path:"/role/{id}" method:"delete" tags:"角色服务" summary:"删除角色"`
	ID     uint `p:"id" v:"required#缺少角色ID" in:"path"`
}

type DeleteRoleRes struct {
}

type GetOneRoleReq struct {
	g.Meta `path:"/role/{id}" method:"get" tags:"角色服务" summary:"获取角色"`
	ID     uint `p:"id" v:"required#缺少角色ID" in:"path"`
}

type GetOneRoleRes struct {
	entity.Role
}

type GetListRoleReq struct {
	g.Meta `path:"/role" method:"get" tags:"角色服务" summary:"通过分页获取角色"`
	api.CommonPaginationReq
}

type GetListRoleRes struct {
	api.CommonPaginationRes
	List []GetOneRoleRes `json:"list" dc:"角色列表数据"`
}

type GetAllRoleReq struct {
	g.Meta `path:"/role/all" method:"get" tags:"角色服务" summary:"获取全部角色"`
}

type GetAllRoleRes struct {
	List []GetOneRoleRes `json:"list" dc:"角色列表数据"`
}
