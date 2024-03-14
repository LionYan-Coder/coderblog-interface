package article

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "coderblog-interface/api/article/v1"
)

func (c *ControllerV1) UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (res *v1.UpdateArticleRes, err error) {
	ctxUser := model.ContextUser{}
	err = g.RequestFromCtx(ctx).GetParam(consts.JWT_PAYLOAD).Scan(&ctxUser)
	if err != nil {
		return
	}
	_, err = service.Article().Update(ctx, model.ArticleUpdateInput{
		Id:      req.Id,
		Title:   req.Title,
		Summary: req.Summary,
		Content: req.Content,
		Author:  ctxUser.Nickname,
	})
	if err != nil {
		return
	}
	return &v1.UpdateArticleRes{}, err
}
