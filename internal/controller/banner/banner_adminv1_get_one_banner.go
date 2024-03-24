package banner

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	adminv1 "coderblog-interface/api/banner/adminV1"
)

func (c *ControllerAdminV1) GetOneBanner(ctx context.Context, req *adminv1.GetOneBannerReq) (res *adminv1.GetOneBannerRes, err error) {
	out, err := service.Banner().GetOne(ctx, model.BannerDetailInput{ID: req.ID})
	if err != nil || out == nil {
		return
	}
	if err = gconv.Scan(out, &res); err != nil {
		return
	}
	return
}
