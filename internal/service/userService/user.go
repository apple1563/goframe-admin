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

func DeleteUser(ctx context.Context, req *vuser.DeleteUserReq) (res *vuser.DeleteUserRes, err error) {
	_, err = dao.User.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func UpdateUser(ctx context.Context, req *vuser.UpdateUserReq) (res *vuser.UpdateUserRes, err error) {
	_, err = dao.User.Ctx(ctx).Where(userCols.Id, req.Id).Data(g.Map{
		userCols.Password: req.Password,
		userCols.Nickname: req.Nickname,
		userCols.Email:    req.Email,
		userCols.Phone:    req.Phone,
		userCols.Status:   req.Status,
		userCols.RoleId:   req.RoleId,
		userCols.RoleName: req.RoleName,
	}).Update()
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

func OneUser(ctx context.Context, req *vuser.OneUserReq) (res *vuser.OneUserRes, err error) {
	err = dao.User.Ctx(ctx).Where(userCols.Id, req.Id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}

func UserInfo(ctx context.Context, req *vuser.UserInfoReq) (res *vuser.UserInfoRes, err error) {
	err = gconv.Scan(ctx.Value("userInfo"), &res)
	if err != nil {
		return nil, err
	}
	return
}

func TreeListUser(ctx context.Context, req *vuser.TreeListUserReq) (res *vuser.TreeListUserRes, err error) {
	var allUsers = make([]*entity.User, 0)
	err = dao.User.Ctx(ctx).Scan(&allUsers)
	if err != nil {
		return nil, err
	}
	var result = &vuser.TreeListUserRes{
		List: genTreeList(allUsers, 0),
	}
	return result, nil
}

// tool
func genTreeList(list []*entity.User, pid int) []*vuser.TreeNodeUser {
	res := make([]*vuser.TreeNodeUser, 0)
	for _, user := range list {
		if gconv.Int(user.Id) == pid {
			var u = &vuser.TreeNodeUser{
				User: user,
			}
			u.Children = genTreeList(list, gconv.Int(user.Id))
			res = append(res, u)
		}
	}
	return res
}
