package article

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "coderblog-interface/api/article/v1"
)

func (c *ControllerV1) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (res *v1.CreateArticleRes, err error) {
	ctxUser := model.ContextUser{}
	err = g.RequestFromCtx(ctx).GetParam(consts.JWT_PAYLOAD).Scan(&ctxUser)
	if err != nil {
		return
	}
	out, err := service.Article().Create(ctx, model.ArticleCreateInput{
		Title:   req.Title,
		Summary: req.Summary,
		Content: req.Content,
		Author:  ctxUser.Nickname,
	})
	if err != nil {
		return
	}
	return &v1.CreateArticleRes{Id: out.Id}, nil
}
