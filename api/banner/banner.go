// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package banner

import (
	"context"

	"coderblog-interface/api/banner/v1"
)

type IBannerV1 interface {
	CreateBanner(ctx context.Context, req *v1.CreateBannerReq) (res *v1.CreateBannerRes, err error)
	UpdateBanner(ctx context.Context, req *v1.UpdateBannerReq) (res *v1.UpdateBannerRes, err error)
	DeleteBanner(ctx context.Context, req *v1.DeleteBannerReq) (res *v1.DeleteBannerRes, err error)
	GetOneBanner(ctx context.Context, req *v1.GetOneBannerReq) (res *v1.GetOneBannerRes, err error)
	GetListBanner(ctx context.Context, req *v1.GetListBannerReq) (res *v1.GetListBannerRes, err error)
	GetListAllBanner(ctx context.Context, req *v1.GetListAllBannerReq) (res *v1.GetListAllBannerRes, err error)
}
