package comment

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/comment/v1"
)

func (c *ControllerV1) GetListComment(ctx context.Context, req *v1.GetListCommentReq) (res *v1.GetListCommentRes, err error) {
	out, err := service.Comment().GetList(ctx, model.CommentListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return
	}
	var list []v1.GetOneCommentRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}

	return &v1.GetListCommentRes{List: list, Total: out.Total, Page: req.Page, Size: req.Size}, nil
}
