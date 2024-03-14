package role

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/role/v1"
)

func (c *ControllerV1) CreateRole(ctx context.Context, req *v1.CreateRoleReq) (res *v1.CreateRoleRes, err error) {
	data := model.RoleCreateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Role().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.CreateRoleRes{ID: out.ID}, nil
}
