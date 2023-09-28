// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:role, do:true"`
	Id        interface{} //
	Status    interface{} // 状态;1:正常2:禁用
	ListOrder interface{} // 排序
	Name      interface{} // 角色名称
	Remark    interface{} // 备注
	DataScope interface{} // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
