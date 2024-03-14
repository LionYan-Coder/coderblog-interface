package menu

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/menu/v1"
)

func (c *ControllerV1) GetListMenu(ctx context.Context, req *v1.GetListMenuReq) (res *v1.GetListMenuRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
