package role

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/role/v1"
)

func (c *ControllerV1) GetAllRole(ctx context.Context, _ *v1.GetAllRoleReq) (res *v1.GetAllRoleRes, err error) {
	out, err := service.Role().GetAll(ctx, model.RoleGetListAllInput{})
	if err != nil {
		return
	}
	var list []v1.GetOneRoleRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return nil, err
	}
	return &v1.GetAllRoleRes{List: list, Total: out.Total}, nil
}
