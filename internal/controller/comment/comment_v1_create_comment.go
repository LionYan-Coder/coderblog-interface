package comment

import (
	"coderblog-interface/internal/consts"
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"

	v1 "coderblog-interface/api/comment/v1"
)

func (c *ControllerV1) CreateComment(ctx context.Context, req *v1.CreateCommentReq) (res *v1.CreateCommentRes, err error) {
	userID := g.RequestFromCtx(ctx).GetParam(consts.AuthIdentifier).Int()
	if err != nil {
		return
	}
	input := model.CommentCreateInput{
		UserID: userID,
	}
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	out, err := service.Comment().Create(ctx, input)
	if err != nil {
		return
	}
	return &v1.CreateCommentRes{ID: out.ID}, nil
}
