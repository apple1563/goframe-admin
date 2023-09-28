// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        uint        `json:"id"        description:""`
	Status    uint        `json:"status"    description:"状态;1:正常2:禁用"`
	ListOrder uint        `json:"listOrder" description:"排序"`
	Name      string      `json:"name"      description:"角色名称"`
	Remark    string      `json:"remark"    description:"备注"`
	DataScope uint        `json:"dataScope" description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
