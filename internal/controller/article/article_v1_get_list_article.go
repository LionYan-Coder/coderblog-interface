package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/article/v1"
)

func (c *ControllerV1) GetListArticle(ctx context.Context, req *v1.GetListArticleReq) (res *v1.GetListArticleRes, err error) {
	out, err := service.Article().GetListByPublish(ctx, model.ArticleListInput{Page: req.Page, Size: req.Size})
	if err != nil {
		return
	}
	var list []v1.ArticleDTO
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	for i, _ := range list {
		list[i].Author, err = utility.GetUserFullName(ctx, out.List[i].UserID)
		if err != nil {
			return
		}
	}
	return &v1.GetListArticleRes{
		List:  list,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total,
	}, nil
}
