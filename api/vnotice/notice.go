package vnotice

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/model/entity"
)

type AddNoticeReq struct {
	g.Meta `tags:"通告" method:"post" path:"/notice"  sm:"添加" dc:"添加通告" `
	*entity.Notice
}
type AddNoticeRes struct{}

type DeleteNoticeReq struct {
	g.Meta `tags:"通告" method:"delete" path:"/notice"  sm:"删除" dc:"删除通告" `
	Id     uint `json:"id" v:"required" description:""`
}
type DeleteNoticeRes struct{}

type UpdateNoticeReq struct {
	g.Meta `tags:"通告" method:"put" path:"/notice"  sm:"修改" dc:"更新通告" `
	*entity.Notice
}
type UpdateNoticeRes struct{}

type ListNoticeReq struct {
	g.Meta `tags:"通告" method:"get" path:"/notice/list"  sm:"列表" dc:"通告列表" `
	*entity.Notice
	*vcommon.CommonPageReq
}
type ListNoticeRes struct {
	List []*entity.Notice `json:"list"`
	*vcommon.CommonPageRes
}

type OneNoticeReq struct {
	g.Meta `tags:"通告" method:"get" path:"/notice"  sm:"单个" dc:"通告" `
	Id     uint `json:"id" v:"required" description:""`
}
type OneNoticeRes *entity.Notice

type ListNoticeForReceiverReq struct {
	g.Meta `tags:"通告" method:"get" path:"/notice/receive/list"  sm:"列表" dc:"通告列表" `
	*entity.Notice
	Status int `json:"status"    description:"1未读2已读3隐藏，用户看过把status置为2"`
	*vcommon.CommonPageReq
}

type ItemNoticeForReceiver struct {
	*entity.Notice
	Uid    uint `json:"uid"       description:""`
	Status int  `json:"status"    `
}

type ListNoticeForReceiverRes struct {
	List []*ItemNoticeForReceiver `json:"list"`
	*vcommon.CommonPageRes
}

type DeleteNoticeForReceiverReq struct {
	g.Meta `tags:"通告" method:"delete" path:"/notice/receive"  sm:"删除" dc:"删除接收通告" `
	Id     uint `json:"id" v:"required" description:""`
}
type DeleteNoticeForReceiverRes struct{}

type UpdateNoticeForReceiverReq struct {
	g.Meta `tags:"通告" method:"put" path:"/notice/receive"  sm:"修改" dc:"设为已读" `
	Status int  `json:"status"   v:"required" description:"1未读2已读3隐藏，用户看过把status置为2"`
	Id     uint `json:"id" v:"required" description:""`
}
type UpdateNoticeForReceiverRes struct{}

type GetNoticeUnreadCountForReceiverReq struct {
	g.Meta `tags:"通告" method:"get" path:"/notice/receive/unread"  sm:"获取" dc:"获取未读消息个数" `
}

type GetNoticeUnreadCountForReceiverRes struct {
	Count int `json:"count"`
}
