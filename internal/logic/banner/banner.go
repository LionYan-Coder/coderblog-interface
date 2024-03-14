package banner

import (
	"coderblog-interface/internal/dao"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"
)

type sBanner struct {
}

func init() {
	service.RegisterBanner(New())
}

func New() service.IBanner {
	return &sBanner{}
}

func (a sBanner) Create(ctx context.Context, in model.BannerCreateInput) (out *model.BannerCreateOutput, err error) {
	id, err := dao.Banner.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.BannerCreateOutput{ID: int(id)}, err
}

func (a sBanner) Update(ctx context.Context, in model.BannerUpdateInput) (out *model.BannerUpdateOutput, err error) {
	_, err = dao.Banner.Ctx(ctx).Data(in).OmitEmpty().Where(dao.Banner.Columns().Id, in.ID).Update()
	return
}

func (a sBanner) Delete(ctx context.Context, in model.BannerDeleteInput) (out *model.BannerDeleteOutput, err error) {
	_, err = dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().Id, in.ID).Delete()
	return
}

func (a sBanner) GetOne(ctx context.Context, in model.BannerDetailInput) (out *model.BannerDetailOutput, err error) {
	err = dao.Banner.Ctx(ctx).Where(dao.Banner.Columns().Id, in.ID).Scan(&out)
	return
}

func (a sBanner) GetAll(ctx context.Context, _ model.BannerListAllInput) (out *model.BannerListAllOutput, err error) {
	m := dao.Banner.Ctx(ctx).OrderDesc(dao.Banner.Columns().UpdateAt)
	out = &model.BannerListAllOutput{}
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	if err = m.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (a sBanner) GetList(ctx context.Context, in model.BannerListInput) (out *model.BannerListOutput, err error) {
	m := dao.Banner.Ctx(ctx).OrderDesc(dao.Banner.Columns().UpdateAt)
	out = &model.BannerListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	listModel := m.Page(in.Page, in.Size)
	total, err := listModel.Count()
	if err != nil || total == 0 {
		return out, err
	}
	out.Total = total
	out.List = make([]model.BannerDetailOutput, 0, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
