package model

import (
	"coderblog-interface/internal/model/entity"
)

type RoleBase struct {
	Name    string `json:"name" description:"角色名称"`
	Type    int    `json:"type" description:"角色类型： 0(管理员) 1（用户）"`
	Comment string `json:"comment" description:"角色注释"`
}

type RoleCreateInput struct {
	RoleBase
}
type RoleCreateOutput struct {
	ID int `json:"id" dc:"id"`
}

type RoleUpdateInput struct {
	ID int `json:"id" dc:"id"`
	RoleBase
}
type RoleUpdateOutput struct {
}
type RoleDeleteInput struct {
	ID int `json:"id" dc:"id"`
}
type RoleDeleteOutput struct {
}

type RoleGetOneInput struct {
	ID int `json:"id" dc:"id"`
}
type RoleGetOneOutput struct {
	entity.Role
}
type RoleGetListInput struct {
	Page int
	Size int
}
type RoleGetListOutput struct {
	Total int                  `json:"total" dc:"总数"`
	Page  int                  `json:"page" dc:"分页号码"`
	Size  int                  `json:"size" dc:"分页数量"`
	List  []RoleListOutputItem `json:"list" description:"角色列表数据"`
}

type RoleListOutputItem struct {
	entity.Role
}
