package notebook

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"coderblog-interface/api/notebook/adminV1"
)

func (c *ControllerAdminV1) GetListNotebook(ctx context.Context, req *adminV1.GetListNotebookReq) (res *adminV1.GetListNotebookRes, err error) {
	out, err := service.Notebook().GetList(ctx, model.NotebookListInput{Page: req.Page, Size: req.Size})
	if err != nil {
		return
	}
	var list []adminV1.GetOneNotebookRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	for i, notebook := range list {
		list[i].Author, err = utility.GetUserFullName(ctx, notebook.UserID)
		if err != nil {
			return
		}
	}
	return &adminV1.GetListNotebookRes{
		List:  list,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total,
	}, nil
}
