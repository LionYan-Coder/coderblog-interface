package permission

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/permission/v1"
)

func (c *ControllerV1) GetOnePermission(ctx context.Context, req *v1.GetOnePermissionReq) (res *v1.GetOnePermissionRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
