package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/model/entity"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/article/v1"
)

func (c *ControllerV1) GetRecentArticleByCurrentMonth(ctx context.Context, req *v1.GetRecentArticleByCurrentMonthReq) (res *v1.GetRecentArticleByCurrentMonthRes, err error) {
	out, err := service.Article().GetRecentByCurrentMonth(ctx, model.ArticleGetRecentByCurrentMonthInput{})
	if err != nil {
		return
	}
	var list []*entity.Article
	if err = gconv.Scan(out.List, &list); err != nil {
		return nil, err
	}
	return &v1.GetRecentArticleByCurrentMonthRes{
		List:  list,
		Total: out.Total,
	}, nil
}
