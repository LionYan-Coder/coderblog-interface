package banner

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	adminv1 "coderblog-interface/api/banner/adminV1"
)

func (c *ControllerAdminV1) GetListBanner(ctx context.Context, req *adminv1.GetListBannerReq) (res *adminv1.GetListBannerRes, err error) {
	out, err := service.Banner().GetList(ctx, model.BannerListInput{Size: req.Size, Page: req.Page})
	if err != nil {
		return
	}
	var list []adminv1.GetOneBannerRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	return &adminv1.GetListBannerRes{List: list, Page: out.Page, Total: out.Total, Size: out.Size}, nil
}
