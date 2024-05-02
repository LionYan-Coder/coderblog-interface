package notebook

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"coderblog-interface/utility"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"coderblog-interface/api/notebook/adminV1"
)

func (c *ControllerAdminV1) CreateNotebook(ctx context.Context, req *adminV1.CreateNotebookReq) (res *adminV1.CreateNotebookRes, err error) {
	var input model.NotebookCreateInput
	if err = gconv.Struct(req, &input); err != nil {
		return
	}
	ctxUser, err := utility.GetUserByHeader(ctx)
	if err != nil {
		return
	}
	input.UserID = ctxUser.ID
	out, err := service.Notebook().Create(ctx, input)
	if err != nil {
		return
	}
	return &adminV1.CreateNotebookRes{
		ID: out.ID,
	}, nil
}
