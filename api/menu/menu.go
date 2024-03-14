// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package menu

import (
	"context"

	"coderblog-interface/api/menu/v1"
)

type IMenuV1 interface {
	CreateMenu(ctx context.Context, req *v1.CreateMenuReq) (res *v1.CreateMenuRes, err error)
	UpdateMenu(ctx context.Context, req *v1.UpdateMenuReq) (res *v1.UpdateMenuRes, err error)
	DeleteMenu(ctx context.Context, req *v1.DeleteMenuReq) (res *v1.DeleteMenuRes, err error)
	GetOneMenu(ctx context.Context, req *v1.GetOneMenuReq) (res *v1.GetOneMenuRes, err error)
	GetListMenu(ctx context.Context, req *v1.GetListMenuReq) (res *v1.GetListMenuRes, err error)
	GetListAllMenu(ctx context.Context, req *v1.GetListAllMenuReq) (res *v1.GetListAllMenuRes, err error)
}
