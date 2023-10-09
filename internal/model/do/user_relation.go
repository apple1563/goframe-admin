// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserRelation is the golang structure of table user_relation for DAO operations like Where/Data.
type UserRelation struct {
	g.Meta  `orm:"table:user_relation, do:true"`
	Id      interface{} //
	PUserId interface{} //
	Level   interface{} //
	UserId  interface{} //
}
