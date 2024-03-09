package user

import (
	"coderblog-interface/internal/service"
	"context"

	v1 "coderblog-interface/api/user/v1"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	res, err = service.User().SignIn(ctx, req)
	return
}
