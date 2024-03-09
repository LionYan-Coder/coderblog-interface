package user

import (
	v1 "coderblog-interface/api/user/v1"
	"coderblog-interface/internal/dao"
	"coderblog-interface/internal/model/do"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {

	return &sUser{}
}

func (s *sUser) SignUp(ctx context.Context, input *v1.SignUpReq) (err error) {

	var (
		available bool
		encrypt   []byte
	)

	available, err = s.IsUsernameAvailable(ctx, input.Username)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`已经存在此用户名"%s"`, input.Username)
	}
	available, err = s.IsNicknameAvailable(ctx, input.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`已经存在此昵称"%s"`, input.Nickname)
	}
	encrypt, err = utility.HashPassword(input.Password)
	if err != nil {
		return err
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, _ gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{Username: input.Username, Nickname: input.Nickname, Password: string(encrypt)}).Insert()
		return err
	})
}

func (s *sUser) SignIn(ctx context.Context, _ *v1.SignInReq) (res *v1.SignInRes, err error) {
	token, expire := service.Auth().LoginHandler(ctx)
	res = &v1.SignInRes{Token: token, Expire: expire}
	return
}

func (s *sUser) IsUsernameAvailable(ctx context.Context, username string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Username: username}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Nickname: nickname}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
