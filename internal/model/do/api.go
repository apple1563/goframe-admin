// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Api is the golang structure of table api for DAO operations like Where/Data.
type Api struct {
	g.Meta    `orm:"table:api, do:true"`
	Id        interface{} //
	Method    interface{} //
	Group     interface{} //
	Url       interface{} //
	Remark    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
