package roleService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vrole"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
)

var roleCols = dao.Role.Columns()

func AddRole(ctx context.Context, req *vrole.AddRoleReq) (res *vrole.AddRoleRes, err error) {
	count, err := dao.Role.Ctx(ctx).Where(roleCols.Name, req.Name).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, consts.ErrRolenameExists
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
	_, err = dao.Role.Ctx(ctx).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func ListRole(ctx context.Context, req *vrole.ListRoleReq) (res *vrole.ListRoleRes, err error) {
	var resp = &vrole.ListRoleRes{}
	resp.List = make([]*entity.Role, 0)
	resp.Size = req.Size
	resp.Page = req.Page
	err = dao.Role.Ctx(ctx).Where(g.Map{
		roleCols.Name + " like ?": req.Name,
		roleCols.Id:               req.Id,
	}).Page(req.Page, req.Size).ScanAndCount(&resp.List, &resp.Total, false)
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
