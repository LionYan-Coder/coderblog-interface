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

func (c *ControllerAdminV1) CreateBanner(ctx context.Context, req *adminv1.CreateBannerReq) (res *adminv1.CreateBannerRes, err error) {
	userID := g.RequestFromCtx(ctx).GetParam(consts.AuthIdentifier).Int()
	input := model.BannerCreateInput{
		UserID: userID,
	}
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	out, err := service.Banner().Create(ctx, input)
	if err != nil {
		return
	}
	return &adminv1.CreateBannerRes{ID: out.ID}, nil
}
