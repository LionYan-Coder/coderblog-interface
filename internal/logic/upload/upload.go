package upload

import (
	"coderblog-interface/internal/model"
	"coderblog-interface/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gfile"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sUpload struct {
}

func init() {
	service.RegisterUpload(New())
}

func New() service.IUpload {
	return &sUpload{}
}

func (u sUpload) UploadImage(ctx context.Context, input model.UploadFileInput) (out *model.UploadFileOutput, err error) {
	file := input.File
	if file == nil {
		err = gerror.Newf(`文件解析失败`)
		return
	}
	fileName, err := file.Save(gfile.Pwd()+"/resource/upload/img/", true)
	if err != nil {
		return
	}
	r := g.RequestFromCtx(ctx)
	schema := r.URL.Scheme
	if g.IsEmpty(schema) {
		schema = "http://"
	}
	return &model.UploadFileOutput{
		Name: file.Filename,
		URL:  schema + r.Host + "/upload/img/" + fileName,
	}, nil
}
