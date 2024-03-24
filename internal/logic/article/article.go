package article

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/dao"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"
)

type sArticle struct {
}

func init() {
	service.RegisterArticle(New())
}

func New() service.IArticle {
	return &sArticle{}
}

func (a sArticle) Create(ctx context.Context, in model.ArticleCreateInput) (out *model.ArticleCreateOutput, err error) {
	id, err := dao.Article.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.ArticleCreateOutput{ID: int(id)}, err
}

func (a sArticle) Update(ctx context.Context, in model.ArticleUpdateInput) (out *model.ArticleUpdateOutput, err error) {
	_, err = dao.Article.Ctx(ctx).Data(in).OmitEmpty().Where(dao.Article.Columns().Id, in.ID).Update()
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

func (a sArticle) GetAll(ctx context.Context, _ model.ArticleListAllInput) (out *model.ArticleListAllOutput, err error) {
	ctxUser := model.ContextUser{}
	err = g.RequestFromCtx(ctx).GetParam(consts.JWT_PAYLOAD).Scan(&ctxUser)
	if err != nil {
		return
	}
	m := dao.Article.Ctx(ctx).Where(dao.Article.Columns().Author, ctxUser.Nickname).OrderDesc(dao.Article.Columns().UpdateAt)
	out = &model.ArticleListAllOutput{}
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	if err = m.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (a sArticle) GetList(ctx context.Context, in model.ArticleListInput) (out *model.ArticleListOutput, err error) {
	ctxUser := model.ContextUser{}
	err = g.RequestFromCtx(ctx).GetParam(consts.JWT_PAYLOAD).Scan(&ctxUser)
	if err != nil {
		return
	}
	m := dao.Article.Ctx(ctx).Where(dao.Article.Columns().Author, ctxUser.Nickname).OrderDesc(dao.Article.Columns().UpdateAt)
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

func (a sArticle) GetRecentByCurrentMonth(ctx context.Context, _ model.ArticleGetRecentByCurrentMonthInput) (out *model.ArticleListAllOutput, err error) {
	curMonth := gtime.Now()
	m := dao.Article.Ctx(ctx).WhereLTE(dao.Article.Columns().CreateAt, curMonth).OrderDesc(dao.Article.Columns().UpdateAt).Limit(5)
	out = &model.ArticleListAllOutput{}
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	if err = m.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
