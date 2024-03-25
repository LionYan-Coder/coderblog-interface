package article

import (
	"coderblog-interface/api/article/adminV1"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/article/v1"
)

func (c *ControllerV1) GetListArticle(ctx context.Context, req *v1.GetListArticleReq) (res *v1.GetListArticleRes, err error) {
	out, err := service.Article().GetList(ctx, model.ArticleListInput{Page: req.Page, Size: req.Size})
	if err != nil {
		return
	}
	var list []*adminV1.GetOneArticleRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	return &v1.GetListArticleRes{
		List:  list,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total,
	}, nil
}
