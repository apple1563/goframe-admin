package vapi

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/model/entity"
)

type AddApiReq struct {
	g.Meta `tags:"api" method:"post" path:"/api"  sm:"添加" dc:"添加api" `
	*entity.Api
}
type AddApiRes struct{}

type DeleteApiReq struct {
	g.Meta `tags:"api" method:"delete" path:"/api"  sm:"添加" dc:"删除api" `
	Id     uint `json:"id" v:"required" description:""`
}
type DeleteApiRes struct{}

type UpdateApiReq struct {
	g.Meta `tags:"api" method:"put" path:"/api"  sm:"修改" dc:"更新api" `
	*entity.Api
}
type UpdateApiRes struct{}

type ListApiReq struct {
	g.Meta `tags:"api" method:"get" path:"/api/list"  sm:"列表" dc:"api列表" `
	*entity.Api
	*vcommon.CommonPageReq
}
type ListApiRes struct {
	List []*entity.Api `json:"list"`
	*vcommon.CommonPageRes
}

type OneApiReq struct {
	g.Meta `tags:"api" method:"get" path:"/api"  sm:"单个" dc:"api" `
	Id     uint `json:"id" v:"required" description:""`
}
type OneApiRes *entity.Api

type ApiForRoleReq struct {
	g.Meta `tags:"api" method:"post" path:"/api/role"  sm:"角色绑定api" dc:"为角色绑定api" `
	RoleId uint          `json:"roleId"`
	Apis   []*entity.Api `json:"apis"`
}

type ApiForRoleRes struct{}

type ApiByRoleReq struct {
	g.Meta `tags:"api" method:"get" path:"/api/role"  sm:"角色绑定api" dc:"获取角色绑定的api" `
	RoleId uint `json:"roleId"`
}
type ApiByRole struct {
	Method string `json:"method"    description:""`
	Url    string `json:"url"       description:""`
}
type ApiByRoleRes struct {
	List []*ApiByRole `json:"list"`
}
