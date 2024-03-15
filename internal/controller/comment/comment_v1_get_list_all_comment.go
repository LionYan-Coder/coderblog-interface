package comment

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/comment/v1"
)

func (c *ControllerV1) GetListAllComment(ctx context.Context, _ *v1.GetListAllCommentReq) (res *v1.GetListAllCommentRes, err error) {
	out, err := service.Comment().GetAll(ctx, model.CommentListAllInput{})
	if err != nil {
		return
	}
	var list []v1.GetOneCommentRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	return &v1.GetListAllCommentRes{List: list, Total: out.Total}, nil
}
