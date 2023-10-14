package vdict

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/model/entity"
)

type AddDictReq struct {
	g.Meta `tags:"字典" method:"post" path:"/dict"  sm:"添加" dc:"添加字典" `
	*entity.Dict
}
type AddDictRes struct{}

type DeleteDictReq struct {
	g.Meta `tags:"字典" method:"delete" path:"/dict"  sm:"添加" dc:"删除字典" `
	Id     uint `json:"id" v:"required" description:""`
}
type DeleteDictRes struct{}

type UpdateDictReq struct {
	g.Meta `tags:"字典" method:"put" path:"/dict"  sm:"修改" dc:"更新字典" `
	*entity.Dict
}
type UpdateDictRes struct{}

type ListDictReq struct {
	g.Meta `tags:"字典" method:"get" path:"/dict/list"  sm:"列表" dc:"字典列表" `
	*entity.Dict
	*vcommon.CommonPageReq
}
type ListDictRes struct {
	List []*entity.Dict `json:"list"`
	*vcommon.CommonPageRes
}

type OneDictReq struct {
	g.Meta `tags:"字典" method:"get" path:"/dict"  sm:"单个" dc:"字典" `
	Id     uint `json:"id" v:"required" description:""`
}
type OneDictRes *entity.Dict
