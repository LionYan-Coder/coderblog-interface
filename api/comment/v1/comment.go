package v1

import (
	"coderblog-interface/api"
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CommentBase struct {
	ReplyID string `p:"replyId" dc:"评论的回复id"`
	Content string `p:"content" v:"required#评论内容不能为空" dc:"评论内容"`
}

type CreateCommentReq struct {
	g.Meta `path:"/comment" method:"put" tags:"内容服务" summary:"创建评论"`
	CommentBase
}

type CreateCommentRes struct {
	ID int `json:"id" dc:"id"`
}

type UpdateCommentReq struct {
	g.Meta `path:"/comment/{id}" method:"put,post" tags:"内容服务" summary:"修改评论"`
	ID     int `in:"path" v:"min:1#缺少评论ID" json:"id" dc:"id"`
	CommentBase
}

type UpdateCommentRes struct {
}

type DeleteCommentReq struct {
	g.Meta `path:"/comment/{id}" method:"delete" tags:"内容服务" summary:"删除评论"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type DeleteCommentRes struct {
}

type GetOneCommentReq struct {
	g.Meta `path:"/comment/{id}" method:"get" tags:"内容服务" summary:"获取评论"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type GetOneCommentRes struct {
	*entity.Comment
}

type GetListCommentReq struct {
	g.Meta `path:"/comment" method:"get" tags:"内容服务" summary:"获取评论"`
	api.CommonPaginationReq
}

type GetListCommentRes struct {
	List  []GetOneCommentRes `json:"list" dc:"评论列表"`
	Total int                `json:"total" dc:"总数"`
	Page  int                `json:"page" dc:"分页号码"`
	Size  int                `json:"size" dc:"分页数量"`
}

type GetListAllCommentReq struct {
	g.Meta `path:"/comment/all" method:"get" tags:"内容服务" summary:"获取所有评论"`
}

type GetListAllCommentRes struct {
	List  []GetOneCommentRes `json:"list" dc:"评论列表"`
	Total int                `json:"total" dc:"评论总数"`
}
