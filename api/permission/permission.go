// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package permission

import (
	"context"

	"coderblog-interface/api/permission/v1"
)

type IPermissionV1 interface {
	CreatePermission(ctx context.Context, req *v1.CreatePermissionReq) (res *v1.CreatePermissionRes, err error)
	UpdatePermission(ctx context.Context, req *v1.UpdatePermissionReq) (res *v1.UpdatePermissionRes, err error)
	DeletePermission(ctx context.Context, req *v1.DeletePermissionReq) (res *v1.DeletePermissionRes, err error)
	GetOnePermission(ctx context.Context, req *v1.GetOnePermissionReq) (res *v1.GetOnePermissionRes, err error)
	GetListPermission(ctx context.Context, req *v1.GetListPermissionReq) (res *v1.GetListPermissionRes, err error)
	GetListAllPermission(ctx context.Context, req *v1.GetListAllPermissionReq) (res *v1.GetListAllPermissionRes, err error)
}
