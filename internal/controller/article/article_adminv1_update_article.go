package article

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) UpdateArticle(ctx context.Context, req *adminV1.UpdateArticleReq) (res *adminV1.UpdateArticleRes, err error) {
	ctxUser := model.ContextUser{}
	err = g.RequestFromCtx(ctx).GetParam(consts.JWT_PAYLOAD).Scan(&ctxUser)
	if err != nil {
		return
	}
	_, err = service.Article().Update(ctx, model.ArticleUpdateInput{
		ID:       req.ID,
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
	return &adminV1.UpdateArticleRes{}, err
}
