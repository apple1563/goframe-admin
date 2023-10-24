// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OtpAuth is the golang structure of table otp_auth for DAO operations like Where/Data.
type OtpAuth struct {
	g.Meta   `orm:"table:otp_auth, do:true"`
	Id       interface{} //
	Uid      interface{} //
	Username interface{} //
	Secret   interface{} //
	Status   interface{} // 1开2关
}
