package vrole

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/model/entity"
)

type AddRoleReq struct {
	g.Meta `tags:"角色" method:"post" path:"/role"  sm:"添加" dc:"添加角色" `
	*entity.Role
}
type AddRoleRes struct{}

type DeleteRoleReq struct {
	g.Meta `tags:"角色" method:"delete" path:"/role"  sm:"添加" dc:"删除角色" `
	Id     uint `json:"id" v:"required" description:""`
}
type DeleteRoleRes struct{}

type UpdateRoleReq struct {
	g.Meta `tags:"角色" method:"put" path:"/role"  sm:"添加" dc:"更新角色" `
	*entity.Role
}
type UpdateRoleRes struct{}

type ListRoleReq struct {
	g.Meta `tags:"角色" method:"get" path:"/role/list"  sm:"列表" dc:"角色列表" `
	*entity.Role
	*vcommon.CommonPageReq
}
type ListRoleRes struct {
	List []*entity.Role `json:"list"`
	*vcommon.CommonPageRes
}

type OneRoleReq struct {
	g.Meta `tags:"角色" method:"get" path:"/role"  sm:"单个" dc:"角色" `
	Id     uint `json:"id" v:"required" description:""`
}
type OneRoleRes *entity.Role
