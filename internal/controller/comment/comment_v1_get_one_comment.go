package comment

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/comment/v1"
)

func (c *ControllerV1) GetOneComment(ctx context.Context, req *v1.GetOneCommentReq) (res *v1.GetOneCommentRes, err error) {
	out, err := service.Comment().GetOne(ctx, model.CommentDetailInput{ID: req.ID})
	if err != nil || out == nil {
		return
	}

	if err = gconv.Scan(out, &res); err != nil {
		return
	}
	return
}
