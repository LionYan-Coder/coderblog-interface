package v1

import "github.com/gogf/gf/v2/frame/g"

type GetOneBannerByCurrentMonthReq struct {
	g.Meta `path:"/banner/currentMonth" method:"get" tags:"横幅服务" summary:"获取当月横幅"`
}

type GetOneBannerByCurrentMonthRes struct {
	Title       string `p:"title" json:"title"`
	Info        string `p:"info" json:"info"`
	CoverURL    string `p:"coverUrl" json:"coverURL"`
	DisplayDate string `p:"displayDate" json:"displayDate"`
}
