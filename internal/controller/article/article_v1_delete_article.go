package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	v1 "coderblog-interface/api/article/v1"
)

func (c *ControllerV1) DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (res *v1.DeleteArticleRes, err error) {
	_, err = service.Article().Delete(ctx, model.ArticleDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return
	}
	return &v1.DeleteArticleRes{}, nil
}
