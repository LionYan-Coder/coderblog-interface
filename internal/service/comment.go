// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"coderblog-interface/internal/model"
	"context"
)

type (
	IComment interface {
		Create(ctx context.Context, in model.CommentCreateInput) (out *model.CommentCreateOutput, err error)
		Update(ctx context.Context, in model.CommentUpdateInput) (out *model.CommentUpdateOutput, err error)
		Delete(ctx context.Context, in model.CommentDeleteInput) (out *model.CommentDeleteOutput, err error)
		GetOne(ctx context.Context, in model.CommentDetailInput) (out *model.CommentDetailOutput, err error)
		GetAll(ctx context.Context, _ model.CommentListAllInput) (out *model.CommentListAllOutput, err error)
		GetList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error)
	}
)

var (
	localComment IComment
)

func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

func RegisterComment(i IComment) {
	localComment = i
}
