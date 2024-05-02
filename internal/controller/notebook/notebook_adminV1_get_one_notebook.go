package notebook

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"coderblog-interface/api/notebook/adminV1"
)

func (c *ControllerAdminV1) GetOneNotebook(ctx context.Context, req *adminV1.GetOneNotebookReq) (res *adminV1.GetOneNotebookRes, err error) {
	out, err := service.Notebook().GetOne(ctx, model.NotebookDetailInput{ID: req.ID})
	if err != nil {
		return
	}
	if err = gconv.Scan(out, &res); err != nil {
		return
	}
	res.Author, err = utility.GetUserFullName(ctx, res.UserID)
	if err != nil {
		return
	}
	return
}
