// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "coderblog-interface/api/user/v1"
	"context"
)

type (
	IUser interface {
		SignUp(ctx context.Context, input v1.SignUpReq) (err error)
		SignIn(ctx context.Context, input v1.SignInReq) (err error)
		IsUsernameAvailable(ctx context.Context, username string) (bool, error)
		IsNicknameAvailable(ctx context.Context, nickname string) (bool, error)
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
