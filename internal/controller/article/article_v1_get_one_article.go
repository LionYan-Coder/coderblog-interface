package article

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/article/v1"
)

func (c *ControllerV1) GetOneArticle(ctx context.Context, req *v1.GetOneArticleReq) (res *v1.GetOneArticleRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
