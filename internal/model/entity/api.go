// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Api is the golang structure for table api.
type Api struct {
	Id        int         `json:"id"        description:""`
	Url       string      `json:"url"       description:""`
	Method    string      `json:"method"    description:""`
	Group     string      `json:"group"     description:""`
	Remark    string      `json:"remark"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
