package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	adminV1 "coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) GetListAllArticle(ctx context.Context, _ *adminV1.GetListAllArticleReq) (res *adminV1.GetListAllArticleRes, err error) {
	out, err := service.Article().GetAll(ctx, model.ArticleListAllInput{})
	if err != nil {
		return
	}
	var list []adminV1.GetOneArticleRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return nil, err
	}
	return &adminV1.GetListAllArticleRes{
		List:  list,
		Total: out.Total,
	}, nil
}
