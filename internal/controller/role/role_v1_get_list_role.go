package role

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/role/v1"
)

func (c *ControllerV1) GetListRole(ctx context.Context, req *v1.GetListRoleReq) (res *v1.GetListRoleRes, err error) {
	out, err := service.Role().GetList(ctx, model.RoleGetListInput{Page: req.Page, Size: req.Size})
	if err != nil {
		return nil, err
	}
	var list []v1.GetOneRoleRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return nil, err
	}
	return &v1.GetListRoleRes{
		List:  list,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total,
	}, nil
}
