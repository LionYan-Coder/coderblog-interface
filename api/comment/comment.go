// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package comment

import (
	"context"

	"coderblog-interface/api/comment/v1"
)

type ICommentV1 interface {
	CreateComment(ctx context.Context, req *v1.CreateCommentReq) (res *v1.CreateCommentRes, err error)
	UpdateComment(ctx context.Context, req *v1.UpdateCommentReq) (res *v1.UpdateCommentRes, err error)
	DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (res *v1.DeleteCommentRes, err error)
	GetOneComment(ctx context.Context, req *v1.GetOneCommentReq) (res *v1.GetOneCommentRes, err error)
	GetListComment(ctx context.Context, req *v1.GetListCommentReq) (res *v1.GetListCommentRes, err error)
	GetListAllComment(ctx context.Context, req *v1.GetListAllCommentReq) (res *v1.GetListAllCommentRes, err error)
}
