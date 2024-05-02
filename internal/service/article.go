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
	IArticle interface {
		Publish(ctx context.Context, in model.ArticlePublishInput) (out *model.ArticlePublishOutput, err error)
		UnPublish(ctx context.Context, in model.ArticleUnPublishInput) (out *model.ArticleUnPublishOutput, err error)
		Create(ctx context.Context, in model.ArticleCreateInput) (out *model.ArticleCreateOutput, err error)
		Update(ctx context.Context, in model.ArticleUpdateInput) (out *model.ArticleUpdateOutput, err error)
		Delete(ctx context.Context, in model.ArticleDeleteInput) (out *model.ArticleDeleteOutput, err error)
		GetOne(ctx context.Context, in model.ArticleDetailInput) (out *model.ArticleDetailOutput, err error)
		GetListByPublish(ctx context.Context, in model.ArticleListInput) (out *model.ArticleListOutput, err error)
		GetList(ctx context.Context, in model.ArticleListInput) (out *model.ArticleListOutput, err error)
	}
)

var (
	localArticle IArticle
)

func Article() IArticle {
	if localArticle == nil {
		panic("implement not found for interface IArticle, forgot register?")
	}
	return localArticle
}

func RegisterArticle(i IArticle) {
	localArticle = i
}
