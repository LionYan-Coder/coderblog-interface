// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"coderblog-interface/internal/model"
	"context"
)

type (
	IBanner interface {
		Create(ctx context.Context, in model.BannerCreateInput) (out *model.BannerCreateOutput, err error)
		Update(ctx context.Context, in model.BannerUpdateInput) (out *model.BannerUpdateOutput, err error)
		Delete(ctx context.Context, in model.BannerDeleteInput) (out *model.BannerDeleteOutput, err error)
		GetOne(ctx context.Context, in model.BannerDetailInput) (out *model.BannerDetailOutput, err error)
		GetAll(ctx context.Context, _ model.BannerListAllInput) (out *model.BannerListAllOutput, err error)
		GetList(ctx context.Context, in model.BannerListInput) (out *model.BannerListOutput, err error)
	}
)

var (
	localBanner IBanner
)

func Banner() IBanner {
	if localBanner == nil {
		panic("implement not found for interface IBanner, forgot register?")
	}
	return localBanner
}

func RegisterBanner(i IBanner) {
	localBanner = i
}
