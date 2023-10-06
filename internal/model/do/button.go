// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Button is the golang structure of table button for DAO operations like Where/Data.
type Button struct {
	g.Meta    `orm:"table:button, do:true"`
	Id        interface{} //
	MenuId    interface{} // 按钮所在菜单id
	MenuTitle interface{} // 按钮所在菜单名称
	Name      interface{} // 按钮标识符
	Title     interface{} // 按钮名称
	Remark    interface{} //
}
