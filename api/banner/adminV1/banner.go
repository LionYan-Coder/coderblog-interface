package adminV1

import (
	"coderblog-interface/api"
	"coderblog-interface/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateBannerReq struct {
	g.Meta      `path:"/banner" method:"put" tags:"横幅服务" summary:"创建横幅"`
	Title       string `p:"title" json:"title" v:"required|max-length:30#请填写标题|标题不能超过30个字符" dc:"横幅标题"`
	Info        string `p:"info" json:"info" v:"max-length:120#概要不能超过120个字符" dc:"横幅概要"`
	CoverURL    string `p:"coverUrl" json:"coverUrl" v:"required|url#缺少横幅封面地址|地址不合法" dc:"横幅封面地址"`
	DisplayDate string `p:"displayDate" json:"displayDate" v:"required|date#缺少横幅显示日期|显示日期不合法" dc:"横幅显示日期"`
}

type CreateBannerRes struct {
	ID int `json:"id" dc:"id"`
}

type UpdateBannerReq struct {
	g.Meta      `path:"/banner/{id}" method:"put,post" tags:"横幅服务" summary:"修改横幅"`
	ID          int    `in:"path" v:"min:1#缺少横幅ID" json:"id" dc:"id"`
	Title       string `p:"title" json:"title" v:"|ax-length:30#标题不能超过30个字符" dc:"横幅标题"`
	Info        string `p:"info" json:"info" v:"max-length:120#概要不能超过120个字符" dc:"横幅概要"`
	CoverURL    string `p:"coverUrl" json:"coverURL" v:"url#地址不合法" dc:"横幅封面地址"`
	DisplayDate string `p:"displayDate" json:"displayDate" v:"date显示日期不合法" dc:"横幅显示日期"`
}

type UpdateBannerRes struct {
}

type DeleteBannerReq struct {
	g.Meta `path:"/banner/{id}" method:"delete" tags:"横幅服务" summary:"删除横幅"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type DeleteBannerRes struct {
}

type GetOneBannerReq struct {
	g.Meta `path:"/banner/{id}" method:"get" tags:"横幅服务" summary:"根据ID获取横幅"`
	ID     int `in:"path" json:"id" dc:"id"`
}

type GetOneBannerRes struct {
	*entity.Banner
}

type GetListBannerReq struct {
	g.Meta `path:"/banner" method:"get" tags:"横幅服务" summary:"获取横幅列表"`
	api.CommonPaginationReq
}

type GetListBannerRes struct {
	List  []GetOneBannerRes `json:"list" dc:"横幅列表"`
	Total int               `json:"total" dc:"总数"`
	Page  int               `json:"page" dc:"分页号码"`
	Size  int               `json:"size" dc:"分页数量"`
}

type GetListAllBannerReq struct {
	g.Meta `path:"/banner/all" method:"get" tags:"横幅服务" summary:"获取所有横幅"`
}

type GetListAllBannerRes struct {
	List  []GetOneBannerRes `json:"list" dc:"横幅列表"`
	Total int               `json:"total" dc:"横幅总数"`
}
