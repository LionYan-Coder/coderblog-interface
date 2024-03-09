package user

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	v1 "coderblog-interface/api/user/v1"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	out, err := service.User().SignIn(ctx, model.UserSignInInput{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return &v1.SignInRes{Token: out.Token, Expire: out.Expire}, nil
}
