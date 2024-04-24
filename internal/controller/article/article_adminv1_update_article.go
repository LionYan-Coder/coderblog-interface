package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) UpdateArticle(ctx context.Context, req *adminV1.UpdateArticleReq) (res *adminV1.UpdateArticleRes, err error) {
	user, err := utility.GetUserByHeader(ctx)
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
		Author:   *user.FirstName + *user.LastName,
		AuthorID: user.ID,
	})
	if err != nil {
		return
	}
	return &adminV1.UpdateArticleRes{}, err
}
