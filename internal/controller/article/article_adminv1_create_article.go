package article

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"coderblog-interface/api/article/adminV1"
)

func (c *ControllerAdminV1) CreateArticle(ctx context.Context, req *adminV1.CreateArticleReq) (res *adminV1.CreateArticleRes, err error) {
	user, err := utility.GetUserByHeader(ctx)
	if err != nil {
		return
	}
	out, err := service.Article().Create(ctx, model.ArticleCreateInput{
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
	return &adminV1.CreateArticleRes{ID: out.ID}, nil
}
