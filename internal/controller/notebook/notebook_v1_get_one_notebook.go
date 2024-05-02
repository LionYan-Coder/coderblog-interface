package notebook

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/notebook/v1"
)

func (c *ControllerV1) GetOneNotebook(ctx context.Context, req *v1.GetOneNotebookReq) (res *v1.GetOneNotebookRes, err error) {
	out, err := service.Notebook().GetOne(ctx, model.NotebookDetailInput{ID: req.ID})
	if err != nil {
		return
	}
	if err = gconv.Scan(out, &res); err != nil {
		return
	}
	res.Author, err = utility.GetUserFullName(ctx, out.UserID)
	if err != nil {
		return
	}
	return
}
