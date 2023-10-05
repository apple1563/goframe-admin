package vbutton

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/model/entity"
)

type AddButtonReq struct {
	g.Meta `tags:"按钮" method:"post" path:"/button"  sm:"添加" dc:"添加按钮" `
	*entity.Button
}
type AddButtonRes struct{}

type DeleteButtonReq struct {
	g.Meta `tags:"按钮" method:"delete" path:"/button"  sm:"添加" dc:"删除按钮" `
	Id     uint `json:"id" v:"required" description:""`
}
type DeleteButtonRes struct{}

type UpdateButtonReq struct {
	g.Meta `tags:"按钮" method:"put" path:"/button"  sm:"添加" dc:"更新按钮" `
	*entity.Button
}
type UpdateButtonRes struct{}

type ListButtonReq struct {
	g.Meta `tags:"按钮" method:"get" path:"/button/list"  sm:"列表" dc:"按钮列表" `
	*entity.Button
	*vcommon.CommonPageReq
}
type ListButtonRes struct {
	List []*entity.Button `json:"list"`
	*vcommon.CommonPageRes
}

type OneButtonReq struct {
	g.Meta `tags:"按钮" method:"get" path:"/button"  sm:"单个" dc:"按钮" `
	Id     uint `json:"id" v:"required" description:""`
}
type OneButtonRes *entity.Button
