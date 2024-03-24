package banner

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	adminv1 "coderblog-interface/api/banner/adminV1"
)

func (c *ControllerAdminV1) DeleteBanner(ctx context.Context, req *adminv1.DeleteBannerReq) (res *adminv1.DeleteBannerRes, err error) {
	_, err = service.Banner().Delete(ctx, model.BannerDeleteInput{ID: req.ID})
	if err != nil {
		return
	}
	return &adminv1.DeleteBannerRes{}, nil
}
