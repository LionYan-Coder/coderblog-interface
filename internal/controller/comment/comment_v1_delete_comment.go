package comment

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	v1 "coderblog-interface/api/comment/v1"
)

func (c *ControllerV1) DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (res *v1.DeleteCommentRes, err error) {
	_, err = service.Comment().Delete(ctx, model.CommentDeleteInput{ID: req.ID})
	if err != nil {
		return
	}
	return &v1.DeleteCommentRes{}, nil
}
