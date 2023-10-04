package userService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-starter/api/vuser"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
	"goframe-starter/utility/xpwd"
)

var userCols = dao.User.Columns()

func AddUser(ctx context.Context, req *vuser.AddUserReq) (res *vuser.AddUserRes, err error) {
	count, err := dao.User.Ctx(ctx).Where(userCols.Username, req.Username).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, consts.ErrUsernameExists
	}
	_, err = dao.User.Ctx(ctx).Data(g.Map{
		userCols.Username: req.Username,
		userCols.RoleId:   req.RoleId,
		userCols.Password: xpwd.GenPwd(req.Password),
	}).Insert()
	if err != nil {
		return nil, err
	}
	return
}

func ListUser(ctx context.Context, req *vuser.ListUserReq) (res *vuser.ListUserRes, err error) {
	var resp = &vuser.ListUserRes{}
	resp.List = make([]*entity.User, 0)
	resp.Size = req.Size
	resp.Page = req.Page
	err = dao.User.Ctx(ctx).Where(g.Map{
		userCols.Username + " like ?": req.Username,
		userCols.Id:                   req.Id,
		userCols.Pid:                  req.Pid,
		userCols.PUsername:            req.PUsername,
		userCols.RoleId:               req.RoleId,
	}).Page(req.Page, req.Size).ScanAndCount(&resp.List, &resp.Total, false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UserInfo(ctx context.Context, req *vuser.UserInfoReq) (res *vuser.UserInfoRes, err error) {
	err = gconv.Scan(ctx.Value("userInfo"), &res)
	if err != nil {
		return nil, err
	}
	return
}
