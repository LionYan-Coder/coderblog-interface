// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package banner

import (
	"context"

	"coderblog-interface/api/banner/adminV1"
	"coderblog-interface/api/banner/v1"
)

type IBannerAdminV1 interface {
	CreateBanner(ctx context.Context, req *adminV1.CreateBannerReq) (res *adminV1.CreateBannerRes, err error)
	UpdateBanner(ctx context.Context, req *adminV1.UpdateBannerReq) (res *adminV1.UpdateBannerRes, err error)
	DeleteBanner(ctx context.Context, req *adminV1.DeleteBannerReq) (res *adminV1.DeleteBannerRes, err error)
	GetOneBanner(ctx context.Context, req *adminV1.GetOneBannerReq) (res *adminV1.GetOneBannerRes, err error)
	GetListBanner(ctx context.Context, req *adminV1.GetListBannerReq) (res *adminV1.GetListBannerRes, err error)
	GetListAllBanner(ctx context.Context, req *adminV1.GetListAllBannerReq) (res *adminV1.GetListAllBannerRes, err error)
}

type IBannerV1 interface {
	GetOneBannerByCurrentMonth(ctx context.Context, req *v1.GetOneBannerByCurrentMonthReq) (res *v1.GetOneBannerByCurrentMonthRes, err error)
}
