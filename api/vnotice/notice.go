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
	g.Meta `tags:"通告" method:"delete" path:"/notice"  sm:"添加" dc:"删除通告" `
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
	g.Meta `tags:"通告" method:"get" path:"/notice/list"  sm:"列表" dc:"通告列表" `
	*entity.Notice
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
