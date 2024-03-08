package v1

import "github.com/gogf/gf/v2/frame/g"

type SignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"用户服务" summary:"用户登录"`
	Username string `p:"username" v:"required#请填写账号"`
	Password string `p:"password" v:"required#请填写密码"`
}

type SignInRes struct {
}
