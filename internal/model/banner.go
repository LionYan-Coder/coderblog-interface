package model

import (
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type BannerCreateInput struct {
	UserID      int         `json:"userId"      ` // 用户id
	Title       string      `json:"title"       ` // 横幅标题
	Info        string      `json:"info"        ` // 横幅信息
	CoverURL    string      `json:"coverUrl"    ` // 横幅图片地址
	DisplayDate *gtime.Time `json:"displayDate" ` // 展示日期
}

type BannerCreateOutput struct {
	ID int
}

type BannerUpdateInput struct {
	ID          int
	UserID      int         `json:"userId"      ` // 用户id
	Title       string      `json:"title"       ` // 横幅标题
	Info        string      `json:"info"        ` // 横幅信息
	CoverURL    string      `json:"coverUrl"    ` // 横幅图片地址
	DisplayDate *gtime.Time `json:"displayDate" ` // 展示日期
}
type BannerUpdateOutput struct {
}
type BannerDeleteInput struct {
	ID int
}
type BannerDeleteOutput struct {
}

type BannerDetailInput struct {
	ID int
}
type BannerDetailOutput struct {
	*entity.Banner
}
type BannerListInput struct {
	Page int
	Size int
}
type BannerListOutput struct {
	List  []BannerDetailOutput
	Total int `json:"total" dc:"总数"`
	Page  int `json:"page" dc:"分页号码"`
	Size  int `json:"size" dc:"分页数量"`
}

type BannerListAllInput struct {
}
type BannerListAllOutput struct {
	Total int
	List  []BannerDetailOutput
}
