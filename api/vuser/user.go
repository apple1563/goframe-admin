package vuser

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/model/entity"
)

type AddUserReq struct {
	g.Meta   `tags:"用户" method:"post" path:"/user"  sm:"添加" dc:"添加用户" `
	Username string `v:"required" dc:"账号" json:"username"`
	Password string `v:"required" dc:"密码" json:"password"`
	RoleId   int    `json:"roleId"        description:"1用户2代理3管理"`
}
type AddUserRes struct{}

type ListUserReq struct {
	g.Meta `tags:"用户" method:"get" path:"/user/list"  sm:"列表" dc:"用来列表" `
	*entity.User
	*vcommon.CommonPageReq
}
type ListUserRes struct {
	List []*entity.User `json:"list"`
	*vcommon.CommonPageRes
}

type UserInfoReq struct {
	g.Meta `tags:"用户" method:"get" path:"/user/info"  sm:"用户信息" dc:"获取当前登录用户信息" `
	//Id     uint `json:"id"  v:"required"  description:""`
}
type UserInfoRes entity.User
