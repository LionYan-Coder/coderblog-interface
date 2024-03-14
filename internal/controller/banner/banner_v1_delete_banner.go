package banner

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	v1 "coderblog-interface/api/banner/v1"
)

func (c *ControllerV1) DeleteBanner(ctx context.Context, req *v1.DeleteBannerReq) (res *v1.DeleteBannerRes, err error) {
	_, err = service.Banner().Delete(ctx, model.BannerDeleteInput{ID: req.ID})
	if err != nil {
		return
	}
	return &v1.DeleteBannerRes{}, nil
}
