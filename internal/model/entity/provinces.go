// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Provinces is the golang structure for table provinces.
type Provinces struct {
	Id        int64       `json:"id"        description:"省市区ID"`
	Title     string      `json:"title"     description:"栏目名称"`
	Pinyin    string      `json:"pinyin"    description:"拼音"`
	Lng       string      `json:"lng"       description:"经度"`
	Lat       string      `json:"lat"       description:"纬度"`
	Pid       int64       `json:"pid"       description:"父栏目"`
	Level     int         `json:"level"     description:"关系树等级"`
	Tree      string      `json:"tree"      description:"关系"`
	Sort      int         `json:"sort"      description:"排序"`
	Status    int         `json:"status"    description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
