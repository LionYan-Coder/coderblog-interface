package banner

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/banner/v1"
)

func (c *ControllerV1) GetOneBannerByCurrentMonth(ctx context.Context, _ *v1.GetOneBannerByCurrentMonthReq) (res *v1.GetOneBannerByCurrentMonthRes, err error) {
	out, err := service.Banner().GetOneByCurrentMonth(ctx, model.BannerGetOneByCurrentMonthInput{})
	if err != nil {
		return
	}
	if err = gconv.Scan(out, &res); err != nil {
		return
	}
	return
}
