package api

import "github.com/gogf/gf/v2/frame/g"

type CommonPaginationReq struct {
	g.Meta `summary:"公共分页参数"`
	Page   int `json:"page" in:"query" d:"1"  v:"min:1#分页号码错误"     dc:"分页号码，默认1"`
	Size   int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}

type CommonPaginationRes struct {
	g.Meta `summary:"公共分页返回数据"`
	Total  int `json:"total" dc:"总数"`
	Page   int `json:"page" dc:"分页号码"`
	Size   int `json:"size" dc:"分页数量"`
}
