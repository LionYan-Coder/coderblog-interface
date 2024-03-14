package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/article/v1"
)

func (c *ControllerV1) GetListAllArticle(ctx context.Context, _ *v1.GetListAllArticleReq) (res *v1.GetListAllArticleRes, err error) {
	out, err := service.Article().GetAll(ctx, model.ArticleListAllInput{})
	if err != nil {
		return
	}
	var list []v1.GetOneArticleRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return nil, err
	}
	return &v1.GetListAllArticleRes{
		List:  list,
		Total: out.Total,
	}, nil
}
