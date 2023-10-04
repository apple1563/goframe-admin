package roleService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vcommon"
	"goframe-starter/api/vrole"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
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
	//  todo 角色绑定的权限规则，如casbin也要删除
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
	var model = dao.Role.Ctx(ctx).Where(data)
	if req.Size != 0 {
		resp.Page = req.Page
		resp.Size = req.Size
		model = model.Page(req.Page, req.Size)
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
