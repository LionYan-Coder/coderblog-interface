package role

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/role/v1"
)

func (c *ControllerV1) GetOneRole(ctx context.Context, req *v1.GetOneRoleReq) (res *v1.GetOneRoleRes, err error) {
	out, err := service.Role().GetOne(ctx, model.RoleGetOneInput{
		ID: req.ID,
	})
	if err != nil || out == nil {
		return nil, err
	}
	if err = gconv.Scan(out, &res); err != nil {
		return nil, err
	}
	return res, nil
}
