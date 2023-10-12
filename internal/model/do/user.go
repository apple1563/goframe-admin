// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta      `orm:"table:user, do:true"`
	Id          interface{} //
	Username    interface{} //
	Password    interface{} //
	Nickname    interface{} //
	Email       interface{} //
	Phone       interface{} //
	Status      interface{} // 1正常2禁用3注销
	ClientAgent interface{} // 注册clientAgen头
	Ip          interface{} // IP
	RoleId      interface{} // 1用户3代理2管理
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	PRoleId     interface{} //
	Pid         interface{} //
	PUsername   interface{} //
	RoleName    interface{} //
	PRoleName   interface{} //
}
