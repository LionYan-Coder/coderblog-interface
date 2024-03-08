package v1

import "github.com/gogf/gf/v2/frame/g"

type SignUpReq struct {
	g.Meta        `path:"/user/sign-up" method:"post" tags:"用户服务" summary:"用户注册"`
	Username      string `p:"username" v:"required|passport#请填写账号|字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	Nickname      string `p:"nickname" v:"required|length:6,10#请填写昵称|昵称长度为:{min}到:{max}位"`
	Password      string `p:"password" v:"required|password2#请填写密码|密码是6-18位并且必须包含大小写字母和数字"`
	CheckPassword string `p:"checkPassword" v:"required|same:password,20#请填写确认密码|与密码不一致" dc:"确认密码"`
}

type SignUpRes struct {
}
