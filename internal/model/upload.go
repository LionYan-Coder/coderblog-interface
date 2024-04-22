package model

import "github.com/gogf/gf/v2/net/ghttp"

type UploadFileInput struct {
	File *ghttp.UploadFile `json:"file"`
}

type UploadFileOutput struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
