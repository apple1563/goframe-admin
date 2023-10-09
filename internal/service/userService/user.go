package userService

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-starter/api/vcommon"
	"goframe-starter/api/vuser"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
	"goframe-starter/utility/xpwd"
)

var userCols = dao.User.Columns()
var roleCols = dao.Role.Columns()
var userRelationCols = dao.UserRelation.Columns()

func AddUser(ctx context.Context, req *vuser.AddUserReq) (res *vuser.AddUserRes, err error) {
	count, err := dao.User.Ctx(ctx).Where(userCols.Username, req.Username).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, consts.ErrUsernameExists
	}
	selfInfo, ok := ctx.Value("userInfo").(*entity.User)
	if !ok {
		return nil, consts.ErrTypeAssertion
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var roleId = 1
		if req.RoleId != 0 {
			roleId = req.RoleId
		}
		value, err := dao.Role.Ctx(ctx).TX(tx).WherePri(roleId).Fields(roleCols.Name).Value()
		if err != nil {
			return err
		}
		id, err := dao.User.Ctx(ctx).TX(tx).Data(g.Map{
			userCols.Username:  req.Username,
			userCols.RoleId:    roleId,
			userCols.RoleName:  value.String(),
			userCols.Password:  xpwd.GenPwd(req.Password),
			userCols.PRoleId:   selfInfo.RoleId,
			userCols.PRoleName: selfInfo.RoleName,
			userCols.PUsername: selfInfo.Username,
			userCols.Pid:       selfInfo.Id,
		}).InsertAndGetId()
		if err != nil {
			return err
		}
		err = createUserRelation(ctx, tx, gconv.Uint(id))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

func createUserRelation(ctx context.Context, tx gdb.TX, id uint) error {
	if id == 0 {
		return nil
	}
	var user *entity.User
	err := dao.User.Ctx(ctx).TX(tx).Where(userCols.Id, id).Scan(&user)
	if err != nil {
		return err
	}
	_, err = dao.UserRelation.Ctx(ctx).TX(tx).Data(g.Map{
		userRelationCols.UserId:  id,
		userRelationCols.PUserId: user.Pid,
		userRelationCols.Level:   1,
	}).Insert()
	if err != nil {
		return err
	}
	//  查找父级的所有父级（直属父级，直属父级的父级。。。）
	var parentRelationList = make([]*entity.UserRelation, 0)
	err = dao.UserRelation.Ctx(ctx).TX(tx).Where(userRelationCols.UserId, user.Pid).Scan(&parentRelationList)
	if err != nil {
		return err
	}
	for _, item := range parentRelationList {
		_, err := dao.UserRelation.Ctx(ctx).TX(tx).Data(g.Map{
			userRelationCols.UserId:  id,
			userRelationCols.PUserId: item.PUserId,
			userRelationCols.Level:   item.Level + 1,
		}).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteUser(ctx context.Context, req *vuser.DeleteUserReq) (res *vuser.DeleteUserRes, err error) {
	// 不能删除，因为关联的东西太多，只可以注销
	_, err = dao.User.Ctx(ctx).WherePri(req.Id).Data(userCols.Status, consts.USER_STATUS_DEAD).Update()
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

/*//  无限制版
func ListUserAll(ctx context.Context, req *vuser.ListUserReq) (res *vuser.ListUserRes, err error) {
	var resp = &vuser.ListUserRes{
		List:          make([]*entity.User, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
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
}*/

// ListUser 只查询当前用户的下级
func ListUser(ctx context.Context, req *vuser.ListUserReq) (res *vuser.ListUserRes, err error) {
	var resp = &vuser.ListUserRes{
		List:          make([]*vuser.ItemUser, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
	resp.Size = req.Size
	resp.Page = req.Page
	var data = g.Map{}
	if req.Username != "" {
		data[userCols.Username+" like ?"] = "%" + req.Username + "%"
	}
	if req.PUsername != "" {
		data[userCols.PUsername+" like ?"] = "%" + req.PUsername + "%"
	}
	if req.Id != 0 {
		data[userCols.Id] = req.Id
	}
	if req.RoleId != 0 {
		data[userCols.RoleId] = req.RoleId
	}

	err = g.Model(dao.User.Table()).LeftJoin(dao.UserRelation.Table()+" ur", "user.id=ur.user_id").Where("ur.p_user_id", ctx.Value("uid")).Fields("user.*,ur.level").Where(data).Page(req.Page, req.Size).ScanAndCount(&resp.List, &resp.Total, false)
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

func TreeListUserScope(ctx context.Context, req *vuser.TreeListUserReq) (res *vuser.TreeListUserRes, err error) {
	var allUsers = make([]*entity.User, 0)
	err = dao.User.Ctx(ctx).Scan(&allUsers)
	if err != nil {
		return nil, err
	}
	var pid = gconv.Int(ctx.Value("uid"))
	// 获取当前层级下的所有用户
	var result = &vuser.TreeListUserRes{
		List: genTreeList(allUsers, pid),
	}
	return result, nil
}

// tool

func genTreeList(list []*entity.User, pid int) []*vuser.TreeNodeUser {
	res := make([]*vuser.TreeNodeUser, 0)
	for _, user := range list {
		if gconv.Int(user.Pid) == pid {
			var u = &vuser.TreeNodeUser{
				User: user,
			}
			u.Children = genTreeList(list, gconv.Int(user.Id))
			res = append(res, u)
		}
	}
	return res
}
