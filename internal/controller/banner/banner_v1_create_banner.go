package banner

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/banner/v1"
)

func (c *ControllerV1) CreateBanner(ctx context.Context, req *v1.CreateBannerReq) (res *v1.CreateBannerRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
