package utility

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"

	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gogf/gf/v2/frame/g"
)

const UserKey = "UserId"

func GetUserByHeader(ctx context.Context) (ctxUser *clerk.User, err error) {
	ctxUser, err = user.Get(ctx, g.RequestFromCtx(ctx).Header.Get(UserKey))
	if err != nil {
		return
	}

	return
}

func GetUserFullName(ctx context.Context, userID string) (fullName string, err error) {
	ctxUser, err := user.Get(ctx, userID)
	if err != nil {
		return
	}
	fullName = *ctxUser.FirstName + *ctxUser.LastName
	return
}

func GetUserNicknameByHeader(ctx context.Context) (nickname string, err error) {
	ctxUser, err := user.Get(ctx, g.RequestFromCtx(ctx).Header.Get(UserKey))
	if err != nil {
		return
	}

	nickname = *ctxUser.FirstName + *ctxUser.LastName
	return
}
