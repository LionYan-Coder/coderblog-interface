package role

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	v1 "coderblog-interface/api/role/v1"
)

func (c *ControllerV1) DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error) {
	_, err = service.Role().Delete(ctx, model.RoleDeleteInput{
		ID: req.ID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRoleRes{}, nil
}
