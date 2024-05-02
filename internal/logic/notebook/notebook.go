package notebook

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/dao"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"
)

type sNotebook struct {
}

func init() {
	service.RegisterNotebook(New())
}

func New() service.INotebook {
	return &sNotebook{}
}

func (n sNotebook) Publish(ctx context.Context, in model.NotebookPublishInput) (out *model.NotebookPublishOutput, err error) {
	user, err := utility.GetUserByHeader(ctx)
	if err != nil {
		return
	}
	_, err = dao.Notebook.Ctx(ctx).Data(dao.Notebook.Columns().Published, consts.Publish).WherePri(in.ID).Where(dao.Notebook.Columns().UserId, user.ID).Update()
	return
}

func (n sNotebook) UnPublish(ctx context.Context, in model.NotebookUnPublishInput) (out *model.NotebookUnPublishOutput, err error) {
	user, err := utility.GetUserByHeader(ctx)
	if err != nil {
		return
	}
	_, err = dao.Notebook.Ctx(ctx).Data(dao.Notebook.Columns().Published, consts.UnPublish).WherePri(in.ID).Where(dao.Notebook.Columns().UserId, user.ID).Update()
	return
}

func (n sNotebook) Create(ctx context.Context, in model.NotebookCreateInput) (out *model.NotebookCreateOutput, err error) {
	id, err := dao.Notebook.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.NotebookCreateOutput{ID: int(id)}, err
}

func (n sNotebook) Update(ctx context.Context, in model.NotebookUpdateInput) (out *model.NotebookUpdateOutput, err error) {
	_, err = dao.Notebook.Ctx(ctx).Data(in).OmitNil().Where(dao.Notebook.Columns().Id, in.ID).Update()
	return
}

func (n sNotebook) Delete(ctx context.Context, in model.NotebookDeleteInput) (out *model.NotebookDeleteOutput, err error) {
	_, err = dao.Notebook.Ctx(ctx).Where(dao.Notebook.Columns().Id, in.ID).Delete()
	return
}

func (n sNotebook) GetOne(ctx context.Context, in model.NotebookDetailInput) (out *model.NotebookDetailOutput, err error) {
	err = dao.Notebook.Ctx(ctx).Where(dao.Notebook.Columns().Id, in.ID).Scan(&out)
	return
}

func (n sNotebook) GetList(ctx context.Context, in model.NotebookListInput) (out *model.NotebookListOutput, err error) {
	m := dao.Notebook.Ctx(ctx).FieldsEx(dao.Notebook.Columns().Content).OrderDesc(dao.Notebook.Columns().UpdateAt)
	out = &model.NotebookListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	listModel := m.Page(in.Page, in.Size)
	total, err := listModel.Count()
	if err != nil || total == 0 {
		return out, err
	}
	out.Total = total
	out.List = make([]model.NotebookDetailOutput, 0, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (n sNotebook) GetListByPublish(ctx context.Context, in model.NotebookListInput) (out *model.NotebookListOutput, err error) {
	m := dao.Notebook.Ctx(ctx).Where(dao.Notebook.Columns().Published, consts.Publish).FieldsEx(dao.Notebook.Columns().Content, dao.Notebook.Columns().Published).OrderDesc(dao.Notebook.Columns().UpdateAt)
	out = &model.NotebookListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	listModel := m.Page(in.Page, in.Size)
	total, err := listModel.Count()
	if err != nil || total == 0 {
		return out, err
	}
	out.Total = total
	out.List = make([]model.NotebookDetailOutput, 0, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
