package upload

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"coderblog-interface/api/upload/adminV1"
)

func (c *ControllerAdminV1) FileUpload(ctx context.Context, req *adminV1.FileUploadReq) (res *adminV1.FileUploadRes, err error) {
	out, err := service.Upload().UploadImage(ctx, model.UploadFileInput{
		File: req.File,
	})
	if err != nil {
		return
	}

	return &adminV1.FileUploadRes{URL: out.URL, Name: out.Name}, nil
}
