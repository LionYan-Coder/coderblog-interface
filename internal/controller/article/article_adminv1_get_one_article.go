package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	adminV1 "coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) GetOneArticle(ctx context.Context, req *adminV1.GetOneArticleReq) (res *adminV1.GetOneArticleRes, err error) {
	out, err := service.Article().GetOne(ctx, model.ArticleDetailInput{
		ID: req.ID,
	})
	if err != nil || out == nil {
		return
	}
	if err = gconv.Scan(out, &res); err != nil {
		return
	}
	return
}
