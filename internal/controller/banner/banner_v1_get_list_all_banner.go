package banner

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/banner/v1"
)

func (c *ControllerV1) GetListAllBanner(ctx context.Context, _ *v1.GetListAllBannerReq) (res *v1.GetListAllBannerRes, err error) {
	out, err := service.Banner().GetAll(ctx, model.BannerListAllInput{})
	if err != nil {
		return
	}
	var list []v1.GetOneBannerRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	return &v1.GetListAllBannerRes{List: list, Total: out.Total}, nil
}
