package log

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-starter/api/vcommon"
	"goframe-starter/api/vlog"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
)

type Log struct{}

var Ctrl = new(Log)

func (u *Log) ListLoginLog(ctx context.Context, req *vlog.ListLoginLogReq) (res *vlog.ListLoginLogRes, err error) {
	res = &vlog.ListLoginLogRes{
		List:          make([]*entity.LoginLog, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
	res.Page = req.Page
	res.Size = req.Size
	var cols = dao.LoginLog.Columns()
	var model = dao.LoginLog.Ctx(ctx).OrderDesc(cols.UpdatedAt)
	if req.Username != "" {
		model = model.WhereLike(cols.Username, "%"+req.Username+"%")
	}
	if req.Ip != "" {
		model = model.Where(cols.Ip, req.Ip)
	}
	if req.TimeBetween != nil && len(req.TimeBetween) == 2 {
		model = model.WhereBetween(cols.UpdatedAt, req.TimeBetween[0], req.TimeBetween[1])
	}
	// 超级管理员返回全部用户，其他角色返回自己和自己的下级
	var roleId = gconv.Int(ctx.Value("roleId"))
	var uid = gconv.Int(ctx.Value("uid"))
	if roleId == consts.Role_Root_Code {
	} else {
		array, err := dao.UserRelation.Ctx(ctx).Where(dao.UserRelation.Columns().PUserId, uid).Array()
		if err != nil {
			return nil, err
		}
		var slice []int
		for _, value := range array {
			slice = append(slice, gconv.Int(value))
		}
		slice = append(slice, uid)
		model = model.WhereIn(cols.Uid, slice)
	}
	err = model.Page(req.Page, req.Size).ScanAndCount(&res.List, &res.Total, false)
	if err != nil {
		return nil, err
	}
	return
}
