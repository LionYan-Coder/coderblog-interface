package banner

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/banner/v1"
)

func (c *ControllerV1) GetListBanner(ctx context.Context, req *v1.GetListBannerReq) (res *v1.GetListBannerRes, err error) {
	out, err := service.Banner().GetList(ctx, model.BannerListInput{Size: req.Size, Page: req.Page})
	if err != nil {
		return
	}
	var list []v1.GetOneBannerRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	return &v1.GetListBannerRes{List: list, Page: out.Page, Total: out.Total, Size: out.Size}, nil
}
