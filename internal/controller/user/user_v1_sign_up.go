package user

import (
	"coderblog-interface/internal/service"
	"context"

	v1 "coderblog-interface/api/user/v1"
)

func (c *ControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = service.User().SignUp(ctx, req)
	return
}
