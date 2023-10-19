// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure of table file for DAO operations like Where/Data.
type File struct {
	g.Meta     `orm:"table:file, do:true"`
	Id         interface{} //
	Drive      interface{} // 上传驱动：本地，ucloud
	OriginName interface{} // 文件原始名
	FileUrl    interface{} //
	Size       interface{} // 单位M
	Ext        interface{} // 扩展名
	Uid        interface{} // 上传者用户id
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	Remark     interface{} //
}
