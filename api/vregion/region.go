package vregion

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/internal/model/entity"
)

type ListTreeReq struct {
	g.Meta `tags:"省市列表" method:"get" path:"/region/tree"  sm:"列表" dc:"树形列表" `
}
type TreeItem struct {
	*entity.Provinces
	Children []*TreeItem `json:"children"`
}
type ListTreeRes struct {
	List []*TreeItem `json:"list"`
}
