package user

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/user/v1"
)

func (c *ControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	userSignUpInput := model.UserSignUpInput{}
	if err = gconv.Scan(req, &userSignUpInput); err != nil {
		return
	}
	_, err = service.User().SignUp(ctx, userSignUpInput)
	return
}
