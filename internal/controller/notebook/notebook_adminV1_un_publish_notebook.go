package notebook

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"coderblog-interface/api/notebook/adminV1"
)

func (c *ControllerAdminV1) UnPublishNotebook(ctx context.Context, req *adminV1.UnPublishNotebookReq) (res *adminV1.UnPublishNotebookRes, err error) {
	_, err = service.Notebook().UnPublish(ctx, model.NotebookUnPublishInput{ID: req.ID})
	if err != nil {
		return
	}

	return &adminV1.UnPublishNotebookRes{}, nil
}
