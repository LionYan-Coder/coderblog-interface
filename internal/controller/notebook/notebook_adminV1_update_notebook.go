package notebook

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"coderblog-interface/api/notebook/adminV1"
)

func (c *ControllerAdminV1) UpdateNotebook(ctx context.Context, req *adminV1.UpdateNotebookReq) (res *adminV1.UpdateNotebookRes, err error) {
	var input model.NotebookUpdateInput
	if err = gconv.Scan(req, &input); err != nil {
	}
	ctxUser, err := utility.GetUserByHeader(ctx)
	if err != nil {
		return
	}
	input.UserID = ctxUser.ID
	_, err = service.Notebook().Update(ctx, input)
	if err != nil {
		return
	}
	return &adminV1.UpdateNotebookRes{}, nil
}
