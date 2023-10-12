// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Blacklist is the golang structure for table blacklist.
type Blacklist struct {
	Id        int64       `json:"id"        description:"黑名单ID"`
	Ip        string      `json:"ip"        description:"IP地址"`
	Remark    string      `json:"remark"    description:"备注"`
	Status    int         `json:"status"    description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
