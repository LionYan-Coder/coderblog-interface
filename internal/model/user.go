package model

import (
	"time"
)

type ContextUser struct {
	Id       int    `c:"AuthId"`       // 用户id
	Username string `json:"username" ` // 用户名
	Nickname string `json:"nickname" ` // 用户昵称
	Password string `json:"password" ` // 用户密码
}

type UserSignInInput struct {
	Username string
	Password string
}

type UserSignInOutput struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type UserSignUpInput struct {
	Username      string `json:"username" `     // 用户名
	Password      string `json:"password" `     // 用户密码
	Nickname      string `json:"nickname" `     // 用户昵称
	CheckPassword string `json:"checkPassword"` // 确认密码
}

type UserSignUpOutput struct {
}
