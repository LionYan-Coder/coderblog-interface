package menu

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/menu/v1"
)

func (c *ControllerV1) GetListAllMenu(ctx context.Context, req *v1.GetListAllMenuReq) (res *v1.GetListAllMenuRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
