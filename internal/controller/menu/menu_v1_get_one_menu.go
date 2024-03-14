package menu

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/menu/v1"
)

func (c *ControllerV1) GetOneMenu(ctx context.Context, req *v1.GetOneMenuReq) (res *v1.GetOneMenuRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
