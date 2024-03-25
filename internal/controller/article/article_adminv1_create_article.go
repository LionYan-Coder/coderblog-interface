package article

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) CreateArticle(ctx context.Context, req *adminV1.CreateArticleReq) (res *adminV1.CreateArticleRes, err error) {
	ctxUser := model.ContextUser{}
	err = g.RequestFromCtx(ctx).GetParam(consts.JWT_PAYLOAD).Scan(&ctxUser)
	if err != nil {
		return
	}
	out, err := service.Article().Create(ctx, model.ArticleCreateInput{
		Title:    req.Title,
		Summary:  req.Summary,
		Content:  req.Content,
		CoverURL: req.CoverURL,
		Tags:     req.Tags,
		Author:   ctxUser.Nickname,
	})
	if err != nil {
		return
	}
	return &adminV1.CreateArticleRes{ID: out.ID}, nil
}
