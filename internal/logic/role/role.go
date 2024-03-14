package role

import (
	"coderblog-interface/internal/dao"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"
)

type sRole struct {
}

func init() {
	service.RegisterRole(New())
}

func New() service.IRole {
	return &sRole{}
}

func (r sRole) Create(ctx context.Context, in model.RoleCreateInput) (out *model.RoleCreateOutput, err error) {
	id, err := dao.Role.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.RoleCreateOutput{ID: int(id)}, err
}

func (r sRole) Update(ctx context.Context, in model.RoleUpdateInput) (out *model.RoleUpdateOutput, err error) {
	_, err = dao.Role.Ctx(ctx).Data(in).OmitEmpty().Where(dao.Role.Columns().Id, in.ID).Update()
	return
}

func (r sRole) Delete(ctx context.Context, in model.RoleDeleteInput) (out *model.RoleDeleteOutput, err error) {
	_, err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.ID).Delete()
	return
}

func (r sRole) GetOne(ctx context.Context, in model.RoleGetOneInput) (out *model.RoleGetOneOutput, err error) {
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.ID).Scan(&out)
	return
}

func (r sRole) GetAll(ctx context.Context, _ model.RoleGetListAllInput) (out *model.RoleGetListAllOutput, err error) {
	m := dao.Role.Ctx(ctx).OrderDesc(dao.Role.Columns().UpdateAt)
	out = &model.RoleGetListAllOutput{}
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	if err = m.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (r sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	m := dao.Role.Ctx(ctx).OrderDesc(dao.Role.Columns().UpdateAt)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	listModel := m.Page(in.Page, in.Size)
	total, err := listModel.Count()
	if err != nil || total == 0 {
		return out, err
	}
	out.Total = total
	out.List = make([]model.RoleGetOneOutput, 0, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
