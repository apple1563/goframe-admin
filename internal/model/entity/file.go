// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure for table file.
type File struct {
	Id         int         `json:"id"         description:""`
	Drive      string      `json:"drive"      description:"上传驱动：本地，ucloud"`
	OriginName string      `json:"originName" description:"文件原始名"`
	FileUrl    string      `json:"fileUrl"    description:""`
	Size       int64       `json:"size"       description:"单位M"`
	Ext        string      `json:"ext"        description:"扩展名"`
	Uid        int         `json:"uid"        description:"上传者用户id"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:""`
	Remark     string      `json:"remark"     description:""`
}
