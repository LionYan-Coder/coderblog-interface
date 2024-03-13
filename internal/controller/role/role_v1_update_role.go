package role

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/role/v1"
)

func (c *ControllerV1) UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (res *v1.UpdateRoleRes, err error) {
	data := model.RoleUpdateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	_, err = service.Role().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRoleRes{}, nil
}
