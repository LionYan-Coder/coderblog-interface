// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package article

import (
	"context"

	"coderblog-interface/api/article/v1"
)

type IArticleV1 interface {
	CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (res *v1.CreateArticleRes, err error)
	UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (res *v1.UpdateArticleRes, err error)
	DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (res *v1.DeleteArticleRes, err error)
	GetOneArticle(ctx context.Context, req *v1.GetOneArticleReq) (res *v1.GetOneArticleRes, err error)
	GetListArticle(ctx context.Context, req *v1.GetListArticleReq) (res *v1.GetListArticleRes, err error)
	GetListAllArticle(ctx context.Context, req *v1.GetListAllArticleReq) (res *v1.GetListAllArticleRes, err error)
}
