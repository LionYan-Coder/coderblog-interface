package service

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/dao"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/model/do"
	"context"
	"time"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"

	"golang.org/x/crypto/bcrypt"

	"github.com/gogf/gf/v2/os/genv"

	jwt "github.com/gogf/gf-jwt/v2"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	env := genv.Map()
	secretKey := env["secretKey"]
	authService = jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "coderblog",
		Key:             []byte(secretKey),
		Timeout:         time.Minute * 60 * 24,
		MaxRefresh:      time.Minute * 60 * 24,
		IdentityKey:     consts.AuthIdentifier,
		TokenLookup:     "header: Authorization, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
}

func Authenticator(ctx context.Context) (res interface{}, err error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in model.UserSignInInput
	)
	if err = r.Parse(&in); err != nil {
		return
	}
	var user *model.ContextUser
	err = dao.User.Ctx(ctx).Where(do.User{Username: in.Username}).Scan(&user)
	if err != nil {
		return
	}
	if user == nil {
		err = gerror.New(`不存在此用户!`)
		return
	}
	b := comparePassword(user.Password, in.Password)
	if !b {
		err = gerror.New(`用户密码错误!`)
		return
	}

	return gconv.Map(&user), nil
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

func comparePassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
