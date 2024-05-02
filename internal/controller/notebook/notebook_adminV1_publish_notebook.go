package notebook

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"coderblog-interface/api/notebook/adminV1"
)

func (c *ControllerAdminV1) PublishNotebook(ctx context.Context, req *adminV1.PublishNotebookReq) (res *adminV1.PublishNotebookRes, err error) {
	_, err = service.Notebook().Publish(ctx, model.NotebookPublishInput{ID: req.ID})
	if err != nil {
		return
	}

	return &adminV1.PublishNotebookRes{}, nil
}
