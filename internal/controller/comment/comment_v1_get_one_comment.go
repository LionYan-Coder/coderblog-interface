package comment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/comment/v1"
)

func (c *ControllerV1) GetOneComment(ctx context.Context, req *v1.GetOneCommentReq) (res *v1.GetOneCommentRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
