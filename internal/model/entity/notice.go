// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notice is the golang structure for table notice.
type Notice struct {
	Id        int         `json:"id"        description:""`
	Title     string      `json:"title"     description:""`
	Content   string      `json:"content"   description:""`
	Creater   string      `json:"creater"   description:"创建者username，发送者"`
	Receivers string      `json:"receivers" description:"接收者用户id数组"`
	Sort      int         `json:"sort"      description:""`
	Tag       string      `json:"tag"       description:""`
	Remark    string      `json:"remark"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
