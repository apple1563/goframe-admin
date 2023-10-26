package roleService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-starter/api/vcommon"
	"goframe-starter/api/vrole"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
	"goframe-starter/utility/xcasbin"
)

var roleCols = dao.Role.Columns()
var userCols = dao.User.Columns()

func AddRole(ctx context.Context, req *vrole.AddRoleReq) (res *vrole.AddRoleRes, err error) {
	count, err := dao.Role.Ctx(ctx).Where(roleCols.Name, req.Name).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, consts.ErrRoleNameExists
	}
	_, err = dao.Role.Ctx(ctx).Data(g.Map{
		roleCols.ListOrder: req.ListOrder,
		roleCols.Name:      req.Name,
		roleCols.Remark:    req.Remark,
	}).Insert()
	if err != nil {
		return nil, err
	}
	return
}

func UpdateRole(ctx context.Context, req *vrole.UpdateRoleReq) (res *vrole.UpdateRoleRes, err error) {
	_, err = dao.Role.Ctx(ctx).Where(roleCols.Id, req.Id).Data(g.Map{
		roleCols.ListOrder: req.ListOrder,
		roleCols.Name:      req.Name,
		roleCols.Remark:    req.Remark,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func DeleteRole(ctx context.Context, req *vrole.DeleteRoleReq) (res *vrole.DeleteRoleRes, err error) {
	// 删之前还要判断下这个角色有没有被其他用户绑定
	count, err := dao.User.Ctx(ctx).Where(userCols.RoleId, req.Id).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, consts.ErrRoleDelete
	}
	_, err = dao.Role.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	//  角色绑定的权限规则，如casbin也要删除
	var sub = consts.Role_Menu_Prefix + gconv.String(req.Id)
	_, err = xcasbin.Enforcer.RemoveFilteredPolicy(0, sub)
	if err != nil {
		return nil, err
	}
	//  按钮绑定的权限规则，如casbin也要删除
	var sub2 = consts.Role_Button_Prefix + gconv.String(req.Id)
	_, err = xcasbin.Enforcer.RemoveFilteredPolicy(0, sub2)
	if err != nil {
		return nil, err
	}
	//  api绑定的权限规则，如casbin也要删除
	var sub3 = consts.Role_Api_Prefix + gconv.String(req.Id)
	_, err = xcasbin.Enforcer.RemoveFilteredPolicy(0, sub3)
	if err != nil {
		return nil, err
	}
	return
}

func ListRole(ctx context.Context, req *vrole.ListRoleReq) (res *vrole.ListRoleRes, err error) {
	var resp = &vrole.ListRoleRes{
		List:          make([]*entity.Role, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
	var data = g.Map{}
	if req.Name != "" {
		data[roleCols.Name+" like ?"] = "%" + req.Name + "%"
	}
	if req.Id != 0 {
		data[roleCols.Id] = req.Id
	}
	var model = dao.Role.Ctx(ctx).Where(data).Order(roleCols.ListOrder)
	if req.PageSize != 0 {
		resp.PageIndex = req.PageIndex
		resp.PageSize = req.PageSize
		model = model.Page(req.PageIndex, req.PageSize)
	}
	err = model.ScanAndCount(&resp.List, &resp.Total, false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func OneRole(ctx context.Context, req *vrole.OneRoleReq) (res *vrole.OneRoleRes, err error) {
	err = dao.Role.Ctx(ctx).Where(roleCols.Id, req.Id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}

// ListRoleForSelect 添加用户的角色下拉选择框，管理员新增用户可选任意角色。代理新增用户只能选代理和用户两个角色
func ListRoleForSelect(ctx context.Context, req *vrole.ListRoleForSelectReq) (res *vrole.ListRoleForSelectRes, err error) {
	var resp = &vrole.ListRoleForSelectRes{
		List: make([]*entity.Role, 0),
	}
	var roleId = gconv.Uint(ctx.Value("roleId"))
	//超级管理员
	if roleId == consts.Role_Root_Code {
		err := dao.Role.Ctx(ctx).WhereNot(roleCols.Id, roleId).Scan(&resp.List)
		if err != nil {
			return nil, err
		}
	} else if roleId == consts.Role_Admin_Code {
		// 管理员可以选择管理员，代理，用户
		err := dao.Role.Ctx(ctx).WhereIn(roleCols.Id, g.Slice{1, 3, roleId}).Scan(&resp.List)
		if err != nil {
			return nil, err
		}
	} else {
		//菲超级管理员新增用户只能选自身角色和用户角色，用户角色为1
		err := dao.Role.Ctx(ctx).WhereIn(roleCols.Id, g.Slice{1, roleId}).Scan(&resp.List)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}
