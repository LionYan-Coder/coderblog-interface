// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package article

import (
	"context"

	"coderblog-interface/api/article/adminV1"
	"coderblog-interface/api/article/v1"
)

type IArticleAdminV1 interface {
	CreateArticle(ctx context.Context, req *adminV1.CreateArticleReq) (res *adminV1.CreateArticleRes, err error)
	UpdateArticle(ctx context.Context, req *adminV1.UpdateArticleReq) (res *adminV1.UpdateArticleRes, err error)
	DeleteArticle(ctx context.Context, req *adminV1.DeleteArticleReq) (res *adminV1.DeleteArticleRes, err error)
	GetOneArticle(ctx context.Context, req *adminV1.GetOneArticleReq) (res *adminV1.GetOneArticleRes, err error)
	GetListArticle(ctx context.Context, req *adminV1.GetListArticleReq) (res *adminV1.GetListArticleRes, err error)
	GetListAllArticle(ctx context.Context, req *adminV1.GetListAllArticleReq) (res *adminV1.GetListAllArticleRes, err error)
}

type IArticleV1 interface {
	GetRecentArticleByCurrentMonth(ctx context.Context, req *v1.GetRecentArticleByCurrentMonthReq) (res *v1.GetRecentArticleByCurrentMonthRes, err error)
	GetListArticle(ctx context.Context, req *v1.GetListArticleReq) (res *v1.GetListArticleRes, err error)
}
