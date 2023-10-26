// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeUserRelation is the golang structure for table notice_user_relation.
type NoticeUserRelation struct {
	Id        int         `json:"id"        description:""`
	NoticeId  int         `json:"noticeId"  description:""`
	Uid       uint        `json:"uid"       description:""`
	Status    int         `json:"status"    description:"1未读2已读3隐藏，用户看过把status置为2，看完后选择删除就把status置为3"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
