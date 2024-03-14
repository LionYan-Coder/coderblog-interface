package banner

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/banner/v1"
)

func (c *ControllerV1) GetListAllBanner(ctx context.Context, req *v1.GetListAllBannerReq) (res *v1.GetListAllBannerRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
