// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"coderblog-interface/api/role/v1"
)

type IRoleV1 interface {
	CreateRole(ctx context.Context, req *v1.CreateRoleReq) (res *v1.CreateRoleRes, err error)
	UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (res *v1.UpdateRoleRes, err error)
	DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error)
}
