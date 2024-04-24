package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) UnPublishArticle(ctx context.Context, req *adminV1.UnPublishArticleReq) (res *adminV1.UnPublishArticleRes, err error) {
	_, err = service.Article().UnPublish(ctx, model.ArticleUnPublishInput{
		ID: req.ID,
	})
	if err != nil {
		return
	}
	return &adminV1.UnPublishArticleRes{}, err
}
