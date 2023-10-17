package vlog

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/model/entity"
)

type ListLoginLogReq struct {
	g.Meta      `tags:"loginLog" method:"get" path:"/loginLog/list"  sm:"列表" dc:"loginLog列表" `
	Username    string   `json:"username"    description:""`
	Ip          string   `json:"ip"          description:""`
	TimeBetween []string `json:"timeBetween"`
	*vcommon.CommonPageReq
}
type ListLoginLogRes struct {
	List []*entity.LoginLog `json:"list"`
	*vcommon.CommonPageRes
}
