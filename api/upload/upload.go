// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package upload

import (
	"context"

	"coderblog-interface/api/upload/adminV1"
)

type IUploadAdminV1 interface {
	FileUpload(ctx context.Context, req *adminV1.FileUploadReq) (res *adminV1.FileUploadRes, err error)
}
