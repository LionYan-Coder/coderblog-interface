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

func (c *ControllerV1) UpdateComment(ctx context.Context, req *v1.UpdateCommentReq) (res *v1.UpdateCommentRes, err error) {
	userID := g.RequestFromCtx(ctx).GetParam(consts.AuthIdentifier).Int()
	if err != nil {
		return
	}
	input := model.CommentUpdateInput{
		UserID: userID,
	}
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	_, err = service.Comment().Update(ctx, input)
	if err != nil {
		return
	}
	return &v1.UpdateCommentRes{}, nil
}
