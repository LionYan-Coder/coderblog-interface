package article

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/article/v1"
)

func (c *ControllerV1) DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (res *v1.DeleteArticleRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
