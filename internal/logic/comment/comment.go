package comment

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/dao"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

type sComment struct {
}

func init() {
	service.RegisterComment(New())
}

func New() service.IComment {
	return &sComment{}
}

func (a sComment) Create(ctx context.Context, in model.CommentCreateInput) (out *model.CommentCreateOutput, err error) {
	if in.ReplyID > 0 {
		count, err := dao.Comment.Ctx(ctx).Where(dao.Comment.Columns().Id, in.ReplyID).Count()
		if err != nil {
			return nil, err
		}
		if count == 0 {
			return nil, gerror.Newf("回复id：%d 不存在！", in.ReplyID)
		}
	}
	id, err := dao.Comment.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.CommentCreateOutput{ID: int(id)}, err
}

func (a sComment) Update(ctx context.Context, in model.CommentUpdateInput) (out *model.CommentUpdateOutput, err error) {
	_, err = dao.Comment.Ctx(ctx).Data(in).OmitEmpty().Where(dao.Comment.Columns().Id, in.ID).Update()
	return
}

func (a sComment) Delete(ctx context.Context, in model.CommentDeleteInput) (out *model.CommentDeleteOutput, err error) {
	_, err = dao.Comment.Ctx(ctx).Where(dao.Comment.Columns().Id, in.ID).Delete()
	return
}

func (a sComment) GetOne(ctx context.Context, in model.CommentDetailInput) (out *model.CommentDetailOutput, err error) {
	err = dao.Comment.Ctx(ctx).WithAll().Where(dao.Comment.Columns().Id, in.ID).Scan(&out)
	return
}

func (a sComment) GetAll(ctx context.Context, _ model.CommentListAllInput) (out *model.CommentListAllOutput, err error) {
	userID := g.RequestFromCtx(ctx).GetParam(consts.AuthIdentifier).Int()
	m := dao.Comment.Ctx(ctx).WithAll().Where(dao.Comment.Columns().UserId, userID).OrderDesc(dao.Comment.Columns().UpdateAt)
	out = &model.CommentListAllOutput{}
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	if err = m.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (a sComment) GetList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error) {
	userID := g.RequestFromCtx(ctx).GetParam(consts.AuthIdentifier).Int()
	m := dao.Comment.Ctx(ctx).WithAll().Where(dao.Comment.Columns().UserId, userID).OrderDesc(dao.Comment.Columns().UpdateAt)
	out = &model.CommentListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	listModel := m.Page(in.Page, in.Size)
	total, err := listModel.Count()
	if err != nil || total == 0 {
		return out, err
	}
	out.Total = total
	out.List = make([]model.CommentDetailOutput, 0, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
