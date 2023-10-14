// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Dict is the golang structure for table dict.
type Dict struct {
	Id          uint        `json:"id"          description:"参数主键"`
	ConfigName  string      `json:"configName"  description:"参数名称"`
	ConfigKey   string      `json:"configKey"   description:"参数键名"`
	ConfigValue string      `json:"configValue" description:"参数键值"`
	ConfigType  string      `json:"configType"  description:"字典类型"`
	CreateBy    uint        `json:"createBy"    description:"创建者"`
	UpdateBy    uint        `json:"updateBy"    description:"更新者"`
	Remark      string      `json:"remark"      description:"备注"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
