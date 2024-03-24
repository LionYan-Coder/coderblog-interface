package banner

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	adminv1 "coderblog-interface/api/banner/adminV1"
)

func (c *ControllerAdminV1) UpdateBanner(ctx context.Context, req *adminv1.UpdateBannerReq) (res *adminv1.UpdateBannerRes, err error) {
	userID := g.RequestFromCtx(ctx).GetParam(consts.AuthIdentifier).Int()
	input := model.BannerUpdateInput{
		UserID: userID,
	}
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	_, err = service.Banner().Update(ctx, input)
	if err != nil {
		return
	}
	return &adminv1.UpdateBannerRes{}, nil
}
