package middleware

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2/jwt"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

func (s sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s sMiddleware) ClerkAuth(r *ghttp.Request) {
	sessionToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	_, err := jwt.Verify(r.Context(), &jwt.VerifyParams{
		Token: sessionToken,
	})
	if err != nil {
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{Code: http.StatusUnauthorized, Message: "未授权！", Data: nil})
	}
	_user, err := utility.GetUserByHeader(r.Context())
	if err != nil {
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{Code: http.StatusBadGateway, Message: "服务器错误！", Data: nil})
	}
	user := model.RoleContext{}
	err = json.Unmarshal(_user.PrivateMetadata, &user)
	if err != nil {
		return
	}
	if user.Role != consts.ADMINISTRATOR {
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{Code: http.StatusForbidden, Message: "当前权限拒绝访问！", Data: nil})
	}
	r.Middleware.Next()
}
