// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"coderblog-interface/internal/model"
	"context"
)

type (
	IRole interface {
		Create(ctx context.Context, in model.RoleCreateInput) (out *model.RoleCreateOutput, err error)
		Update(ctx context.Context, in model.RoleUpdateInput) (out *model.RoleUpdateOutput, err error)
		Delete(ctx context.Context, in model.RoleDeleteInput) (out *model.RoleDeleteOutput, err error)
		GetOne(ctx context.Context, in model.RoleGetOneInput) (out *model.RoleGetOneOutput, err error)
		GetAll(ctx context.Context, _ model.RoleGetListAllInput) (out *model.RoleGetListAllOutput, err error)
		GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
