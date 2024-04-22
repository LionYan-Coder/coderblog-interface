package middleware

import (
	"coderblog-interface/internal/service"
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

func (s sMiddleware) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

func (s sMiddleware) ClerkAuth(r *ghttp.Request) {
	sessionToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	// Verify the session
	_, err := jwt.Verify(r.Context(), &jwt.VerifyParams{
		Token: sessionToken,
	})
	if err != nil {
		r.Response.WriteJsonExit(ghttp.DefaultHandlerResponse{Code: http.StatusUnauthorized, Message: "Unauthorized", Data: nil})
	}
	r.Middleware.Next()
}
