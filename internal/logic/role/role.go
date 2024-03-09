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

func (r sRole) Update(ctx context.Context, in model.RoleUpdateInput) (err error) {
	_, err = dao.Role.Ctx(ctx).Data(in).OmitEmpty().Where(dao.Role.Columns().Id, in.ID).Update()
	return
}
