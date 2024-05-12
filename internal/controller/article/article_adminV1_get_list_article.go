package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) GetListArticle(ctx context.Context, req *adminV1.GetListArticleReq) (res *adminV1.GetListArticleRes, err error) {
	out, err := service.Article().GetList(ctx, model.ArticleListInput{Page: req.Page, Size: req.Size})
	if err != nil {
		return
	}
	var list []adminV1.GetOneArticleRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	for i, article := range list {
		list[i].Author, err = utility.GetUserFullName(ctx, article.UserID)
		if err != nil {
			return
		}
	}
	return &adminV1.GetListArticleRes{
		List:  list,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total,
	}, nil
}
