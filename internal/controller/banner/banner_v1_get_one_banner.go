package banner

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/banner/v1"
)

func (c *ControllerV1) GetOneBanner(ctx context.Context, req *v1.GetOneBannerReq) (res *v1.GetOneBannerRes, err error) {
	out, err := service.Banner().GetOne(ctx, model.BannerDetailInput{ID: req.ID})
	if err != nil || out == nil {
		return
	}
	if err = gconv.Scan(out, &res); err != nil {
		return
	}
	return
}
