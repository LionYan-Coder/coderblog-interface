package notebook

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "coderblog-interface/api/notebook/v1"
)

func (c *ControllerV1) GetListNotebook(ctx context.Context, req *v1.GetListNotebookReq) (res *v1.GetListNotebookRes, err error) {
	out, err := service.Notebook().GetListByPublish(ctx, model.NotebookListInput{Size: req.Size, Page: req.Size})
	if err != nil {
		return
	}
	var list []v1.GetOneNotebookRes
	if err = gconv.Scan(out.List, &list); err != nil {
		return
	}
	for i := 0; i < len(out.List); i++ {
		list[i].Author, err = utility.GetUserFullName(ctx, out.List[i].UserID)
		if err != nil {
			return
		}
	}
	return &v1.GetListNotebookRes{
		List:  list,
		Page:  req.Page,
		Size:  req.Size,
		Total: out.Total,
	}, nil
}
