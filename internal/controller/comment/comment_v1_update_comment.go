package comment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/comment/v1"
)

func (c *ControllerV1) UpdateComment(ctx context.Context, req *v1.UpdateCommentReq) (res *v1.UpdateCommentRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
