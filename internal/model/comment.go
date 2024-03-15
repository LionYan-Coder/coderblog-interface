package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type CommentCreateInput struct {
	UserID    int    `json:"userId"`
	ArticleID int    `json:"articleId"`
	ReplyID   int    `json:"replyId"`
	Content   string `json:"content"`
}

type CommentCreateOutput struct {
	ID int `json:"id"`
}

type CommentUpdateInput struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	ArticleID int    `json:"articleId"`
	ReplyID   int    `json:"replyId"`
	Content   string `json:"content"`
}

type CommentUpdateOutput struct {
}

type CommentDetailInput struct {
	ID int `json:"id"`
}

type CommentDetailOutput struct {
	ID        int         `json:"id"        `
	UserID    int         `json:"userId"`
	ArticleID int         `json:"articleId" `
	ReplyID   int         `json:"replyId"   `
	Content   string      `json:"content"   `
	User      *UserBase   `json:"userInfo" orm:"with:id=user_id"`
	CreateAt  *gtime.Time `json:"createAt"  `
	UpdateAt  *gtime.Time `json:"updateAt"  `
}

type CommentDeleteInput struct {
	ID int `json:"id"`
}

type CommentDeleteOutput struct {
}

type CommentListInput struct {
	Page int
	Size int
}

type CommentListOutput struct {
	List  []CommentDetailOutput
	Total int `json:"total" dc:"总数"`
	Page  int `json:"page" dc:"分页号码"`
	Size  int `json:"size" dc:"分页数量"`
}

type CommentListAllInput struct {
}
type CommentListAllOutput struct {
	Total int
	List  []CommentDetailOutput
}
