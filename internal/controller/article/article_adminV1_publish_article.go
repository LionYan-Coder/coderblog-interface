package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) PublishArticle(ctx context.Context, req *adminV1.PublishArticleReq) (res *adminV1.PublishArticleRes, err error) {
	_, err = service.Article().Publish(ctx, model.ArticlePublishInput{
		ID: req.ID,
	})
	if err != nil {
		return
	}
	return &adminV1.PublishArticleRes{}, err
}
