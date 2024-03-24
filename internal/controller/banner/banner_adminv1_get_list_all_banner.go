package banner

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	adminv1 "coderblog-interface/api/banner/adminV1"
)

func (c *ControllerAdminV1) GetListAllBanner(ctx context.Context, _ *adminv1.GetListAllBannerReq) (res *adminv1.GetListAllBannerRes, err error) {
	out, err := service.Banner().GetAll(ctx, model.BannerListAllInput{})
	if err != nil {
		return
	}
	var list []adminv1.GetOneBannerRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	return &adminv1.GetListAllBannerRes{List: list, Total: out.Total}, nil
}
