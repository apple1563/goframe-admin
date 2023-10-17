// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LoginLog is the golang structure of table login_log for DAO operations like Where/Data.
type LoginLog struct {
	g.Meta      `orm:"table:login_log, do:true"`
	Id          interface{} //
	Uid         interface{} //
	Username    interface{} //
	Ip          interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	ClientAgent interface{} // 注册clientAgen头
	RoleId      interface{} // 1用户2代理3管理
	PRoleId     interface{} //
	RoleName    interface{} //
	Pid         interface{} //
	PUsername   interface{} //
}
