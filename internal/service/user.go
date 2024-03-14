// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"coderblog-interface/internal/model"
	"context"
)

type (
	IUser interface {
		SignUp(ctx context.Context, input model.UserSignUpInput) (res *model.UserSignUpOutput, err error)
		SignIn(ctx context.Context, _ model.UserSignInInput) (res *model.UserSignInOutput, err error)
		IsUsernameAvailable(ctx context.Context, username string) (bool, error)
		IsNicknameAvailable(ctx context.Context, nickname string) (bool, error)
		GetCtxUser(ctx context.Context) (res *model.ContextUser, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
