package article

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/dao"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"
)

type sArticle struct {
}

func init() {
	service.RegisterArticle(New())
}

func New() service.IArticle {
	return &sArticle{}
}

func (a sArticle) Publish(ctx context.Context, in model.ArticlePublishInput) (out *model.ArticlePublishOutput, err error) {
	user, err := utility.GetUserByHeader(ctx)
	if err != nil {
		return
	}
	_, err = dao.Article.Ctx(ctx).Data(dao.Article.Columns().Published, consts.Publish).WherePri(in.ID).Where(dao.Article.Columns().UserId, user.ID).Update()
	return
}

func (a sArticle) UnPublish(ctx context.Context, in model.ArticleUnPublishInput) (out *model.ArticleUnPublishOutput, err error) {
	user, err := utility.GetUserByHeader(ctx)
	if err != nil {
		return
	}
	_, err = dao.Article.Ctx(ctx).Data(dao.Article.Columns().Published, consts.UnPublish).WherePri(in.ID).Where(dao.Article.Columns().UserId, user.ID).Update()
	return
}

func (a sArticle) Create(ctx context.Context, in model.ArticleCreateInput) (out *model.ArticleCreateOutput, err error) {
	id, err := dao.Article.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.ArticleCreateOutput{ID: int(id)}, err
}

func (a sArticle) Update(ctx context.Context, in model.ArticleUpdateInput) (out *model.ArticleUpdateOutput, err error) {
	_, err = dao.Article.Ctx(ctx).Data(in).OmitNil().Where(dao.Article.Columns().Id, in.ID).Update()
	return
}

func (a sArticle) Delete(ctx context.Context, in model.ArticleDeleteInput) (out *model.ArticleDeleteOutput, err error) {
	_, err = dao.Article.Ctx(ctx).Where(dao.Article.Columns().Id, in.ID).Delete()
	return
}

func (a sArticle) GetOne(ctx context.Context, in model.ArticleDetailInput) (out *model.ArticleDetailOutput, err error) {
	err = dao.Article.Ctx(ctx).Where(dao.Article.Columns().Id, in.ID).Scan(&out)
	return
}

func (a sArticle) GetList(ctx context.Context, in model.ArticleListInput) (out *model.ArticleListOutput, err error) {
	m := dao.Article.Ctx(ctx).FieldsEx(dao.Article.Columns().Content).OrderDesc(dao.Article.Columns().UpdateAt)
	out = &model.ArticleListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	listModel := m.Page(in.Page, in.Size)
	total, err := listModel.Count()
	if err != nil || total == 0 {
		return out, err
	}
	out.Total = total
	out.List = make([]model.ArticleDetailOutput, 0, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (a sArticle) GetListByPublish(ctx context.Context, in model.ArticleListInput) (out *model.ArticleListOutput, err error) {
	m := dao.Article.Ctx(ctx).Where(dao.Article.Columns().Published, consts.Publish).FieldsEx(dao.Article.Columns().Content, dao.Article.Columns().Published).OrderDesc(dao.Article.Columns().UpdateAt)
	out = &model.ArticleListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	listModel := m.Page(in.Page, in.Size)
	total, err := listModel.Count()
	if err != nil || total == 0 {
		return out, err
	}
	out.Total = total
	out.List = make([]model.ArticleDetailOutput, 0, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
