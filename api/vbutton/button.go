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
	g.Meta `tags:"按钮" method:"put" path:"/button"  sm:"修改" dc:"更新按钮" `
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

type ButtonForRoleReq struct {
	g.Meta    `tags:"按钮" method:"post" path:"/button/role"  sm:"角色绑定按钮" dc:"为角色绑定按钮" `
	RoleId    uint  `json:"roleId"`
	ButtonIds []int `json:"buttonIds"`
}

type ButtonForRoleRes struct{}

type ButtonByRoleReq struct {
	g.Meta `tags:"按钮" method:"get" path:"/button/role"  sm:"角色绑定按钮" dc:"获取角色绑定的按钮" `
	RoleId uint `json:"roleId"`
}

type ButtonByRoleRes struct {
	List []int `json:"list"`
}

type ButtonWhitelistReq struct {
	g.Meta `tags:"按钮" method:"get" path:"/button/whitelist"  sm:"角色绑定按钮" dc:"获取当前角色绑定的按钮" `
}

type ButtonWhitelistRes struct {
	List []int `json:"list"`
}
