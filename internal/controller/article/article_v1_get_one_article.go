package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/article/v1"
)

func (c *ControllerV1) GetOneArticle(ctx context.Context, req *v1.GetOneArticleReq) (res *v1.GetOneArticleRes, err error) {
	out, err := service.Article().GetOne(ctx, model.ArticleDetailInput{
		ID: req.ID,
	})
	if err != nil || out == nil {
		return
	}
	if err = gconv.Scan(out, &res); err != nil {
		return
	}
	res.Author, err = utility.GetUserFullName(ctx, out.UserID)
	if err != nil {
		return
	}
	return
}
