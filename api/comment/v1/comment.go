package v1

import (
	"coderblog-interface/api"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"
)

type CommentBase struct {
	ArticleID int    `p:"articleId" v:"min:1#缺少文章ID" dc:"文章id"`
	ReplyID   int    `p:"replyId" v:"different:id#回复ID不能与当前ID相同"  dc:"评论的回复id"`
	Content   string `p:"content" v:"required#评论内容不能为空" dc:"评论内容"`
}

type CreateCommentReq struct {
	g.Meta `path:"/comment" method:"put" tags:"评论服务" summary:"创建评论"`
	CommentBase
}

type CreateCommentRes struct {
	ID int `json:"id" dc:"id"`
}

type UpdateCommentReq struct {
	g.Meta `path:"/comment/{id}" method:"put,post" tags:"评论服务" summary:"修改评论"`
	ID     int `in:"path" v:"min:1#缺少评论ID" p:"id" dc:"id"`
	CommentBase
}

type UpdateCommentRes struct {
}

type DeleteCommentReq struct {
	g.Meta `path:"/comment/{id}" method:"delete" tags:"评论服务" summary:"删除评论"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type DeleteCommentRes struct {
}

type GetOneCommentReq struct {
	g.Meta `path:"/comment/{id}" method:"get" tags:"评论服务" summary:"根据ID获取评论"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type GetOneCommentRes struct {
	ID        int         `json:"id" dc:"评论id" `
	User      interface{} `json:"userInfo" dc:"用户" `
	ArticleID int         `json:"articleId" dc:"评论的文章id" `
	ReplyID   int         `json:"replyId" dc:"回复评论id"  `
	Content   string      `json:"content" dc:"评论内容"  `
	CreateAt  *gtime.Time `json:"createAt" dc:"创建时间"  `
	UpdateAt  *gtime.Time `json:"updateAt" dc:"更新时间" `
}

type GetListCommentReq struct {
	g.Meta `path:"/comment" method:"get" tags:"评论服务" summary:"获取评论列表"`
	api.CommonPaginationReq
}

type GetListCommentRes struct {
	List  []GetOneCommentRes `json:"list" dc:"评论列表"`
	Total int                `json:"total" dc:"总数"`
	Page  int                `json:"page" dc:"分页号码"`
	Size  int                `json:"size" dc:"分页数量"`
}

type GetListAllCommentReq struct {
	g.Meta `path:"/comment/all" method:"get" tags:"评论服务" summary:"获取所有评论"`
}

type GetListAllCommentRes struct {
	List  []GetOneCommentRes `json:"list" dc:"评论列表"`
	Total int                `json:"total" dc:"评论总数"`
}
