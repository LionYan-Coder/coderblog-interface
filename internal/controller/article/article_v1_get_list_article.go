package article

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/article/v1"
)

func (c *ControllerV1) GetListArticle(ctx context.Context, req *v1.GetListArticleReq) (res *v1.GetListArticleRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
