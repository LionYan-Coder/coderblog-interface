package menu

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/menu/v1"
)

func (c *ControllerV1) CreateMenu(ctx context.Context, req *v1.CreateMenuReq) (res *v1.CreateMenuRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
