// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeUserRelation is the golang structure of table notice_user_relation for DAO operations like Where/Data.
type NoticeUserRelation struct {
	g.Meta    `orm:"table:notice_user_relation, do:true"`
	Id        interface{} //
	NoticeId  interface{} //
	Uid       interface{} //
	Status    interface{} // 1未读2已读3隐藏，用户看过把status置为2，看完后选择删除就把status置为3
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
