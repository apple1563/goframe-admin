// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notice is the golang structure of table notice for DAO operations like Where/Data.
type Notice struct {
	g.Meta    `orm:"table:notice, do:true"`
	Id        interface{} //
	Title     interface{} //
	Content   interface{} //
	Creater   interface{} // 创建者username，发送者
	Receivers interface{} // 接收者用户id数组
	Sort      interface{} //
	Tag       interface{} //
	Remark    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
