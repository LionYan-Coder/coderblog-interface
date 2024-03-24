package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	adminV1 "coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) DeleteArticle(ctx context.Context, req *adminV1.DeleteArticleReq) (res *adminV1.DeleteArticleRes, err error) {
	_, err = service.Article().Delete(ctx, model.ArticleDeleteInput{
		ID: req.ID,
	})
	if err != nil {
		return
	}
	return &adminV1.DeleteArticleRes{}, nil
}
