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
	INotebook interface {
		Publish(ctx context.Context, in model.NotebookPublishInput) (out *model.NotebookPublishOutput, err error)
		UnPublish(ctx context.Context, in model.NotebookUnPublishInput) (out *model.NotebookUnPublishOutput, err error)
		Create(ctx context.Context, in model.NotebookCreateInput) (out *model.NotebookCreateOutput, err error)
		Update(ctx context.Context, in model.NotebookUpdateInput) (out *model.NotebookUpdateOutput, err error)
		Delete(ctx context.Context, in model.NotebookDeleteInput) (out *model.NotebookDeleteOutput, err error)
		GetOne(ctx context.Context, in model.NotebookDetailInput) (out *model.NotebookDetailOutput, err error)
		GetList(ctx context.Context, in model.NotebookListInput) (out *model.NotebookListOutput, err error)
		GetListByPublish(ctx context.Context, in model.NotebookListInput) (out *model.NotebookListOutput, err error)
	}
)

var (
	localNotebook INotebook
)

func Notebook() INotebook {
	if localNotebook == nil {
		panic("implement not found for interface INotebook, forgot register?")
	}
	return localNotebook
}

func RegisterNotebook(i INotebook) {
	localNotebook = i
}
