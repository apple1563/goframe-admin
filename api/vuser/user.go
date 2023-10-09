package vuser

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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
type DeleteUserReq struct {
	g.Meta `tags:"用户" method:"delete" path:"/user"  sm:"添加" dc:"删除用户" `
	Id     uint `json:"id" v:"required" description:""`
}
type DeleteUserRes struct{}

type UpdateUserReq struct {
	g.Meta `tags:"用户" method:"put" path:"/user"  sm:"修改" dc:"更新用户" `
	*entity.User
}
type UpdateUserRes struct{}

type ListUserReq struct {
	g.Meta `tags:"用户" method:"get" path:"/user/list"  sm:"列表" dc:"用来列表" `
	*entity.User
	*vcommon.CommonPageReq
}
type UserInfo struct {
	Id          uint        `json:"id"          description:""`
	Username    string      `json:"username"    description:""`
	Nickname    string      `json:"nickname"    description:""`
	Email       string      `json:"email"       description:""`
	Phone       string      `json:"phone"       description:""`
	Status      int         `json:"status"      description:"1正常2禁用"`
	ClientAgent string      `json:"clientAgent" description:"注册clientAgen头"`
	Ip          string      `json:"ip"          description:"IP"`
	RoleId      int         `json:"roleId"      description:"1用户2代理3管理"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
	PRoleId     int         `json:"pRoleId"     description:""`
	Pid         int         `json:"pid"         description:""`
	PUsername   string      `json:"pUsername"   description:""`
	RoleName    string      `json:"roleName"    description:""`
	PRoleName   string      `json:"pRoleName"   description:""`
}
type ItemUser struct {
	*UserInfo
	Level uint `json:"level"   description:""`
}
type ListUserRes struct {
	List []*ItemUser `json:"list"`
	*vcommon.CommonPageRes
}

type OneUserReq struct {
	g.Meta `tags:"用户" method:"get" path:"/user"  sm:"单个" dc:"用户" `
	Id     uint `json:"id" v:"required" description:""`
}
type OneUserRes *entity.User

type UserInfoReq struct {
	g.Meta `tags:"用户" method:"get" path:"/user/info"  sm:"用户信息" dc:"获取当前登录用户信息" `
	//Id     uint `json:"id"  v:"required"  description:""`
}
type UserInfoRes UserInfo

// TreeListUserReq 上下级树
type TreeListUserReq struct {
	g.Meta `tags:"用户" method:"get" path:"/user/tree"  sm:"单个" dc:"用户树" `
}

type TreeNodeUser struct {
	*entity.User
	Children []*TreeNodeUser `json:"children"`
}

type TreeListUserRes struct {
	List []*TreeNodeUser `json:"list"`
}

type LogoutReq struct {
	g.Meta `tags:"用户" method:"get" path:"/logout"  sm:"登出" dc:"登出" `
}

type LogoutRes struct{}
