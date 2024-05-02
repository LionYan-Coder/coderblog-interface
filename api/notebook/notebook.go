// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package notebook

import (
	"context"

	"coderblog-interface/api/notebook/adminV1"
	"coderblog-interface/api/notebook/v1"
)

type INotebookAdminV1 interface {
	CreateNotebook(ctx context.Context, req *adminV1.CreateNotebookReq) (res *adminV1.CreateNotebookRes, err error)
	UpdateNotebook(ctx context.Context, req *adminV1.UpdateNotebookReq) (res *adminV1.UpdateNotebookRes, err error)
	DeleteNotebook(ctx context.Context, req *adminV1.DeleteNotebookReq) (res *adminV1.DeleteNotebookRes, err error)
	GetOneNotebook(ctx context.Context, req *adminV1.GetOneNotebookReq) (res *adminV1.GetOneNotebookRes, err error)
	GetListNotebook(ctx context.Context, req *adminV1.GetListNotebookReq) (res *adminV1.GetListNotebookRes, err error)
	PublishNotebook(ctx context.Context, req *adminV1.PublishNotebookReq) (res *adminV1.PublishNotebookRes, err error)
	UnPublishNotebook(ctx context.Context, req *adminV1.UnPublishNotebookReq) (res *adminV1.UnPublishNotebookRes, err error)
}

type INotebookV1 interface {
	GetListNotebook(ctx context.Context, req *v1.GetListNotebookReq) (res *v1.GetListNotebookRes, err error)
	GetOneNotebook(ctx context.Context, req *v1.GetOneNotebookReq) (res *v1.GetOneNotebookRes, err error)
}
