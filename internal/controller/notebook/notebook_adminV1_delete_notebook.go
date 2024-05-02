package notebook

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"coderblog-interface/api/notebook/adminV1"
)

func (c *ControllerAdminV1) DeleteNotebook(ctx context.Context, req *adminV1.DeleteNotebookReq) (res *adminV1.DeleteNotebookRes, err error) {
	_, err = service.Notebook().Delete(ctx, model.NotebookDeleteInput{ID: req.ID})
	if err != nil {
		return
	}

	return &adminV1.DeleteNotebookRes{}, nil
}
