package buttonService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-starter/api/vbutton"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
	"goframe-starter/utility/xcasbin"
)

var buttonCols = dao.Button.Columns()

func AddButton(ctx context.Context, req *vbutton.AddButtonReq) (res *vbutton.AddButtonRes, err error) {
	count, err := dao.Button.Ctx(ctx).Where(g.Map{
		buttonCols.Name:   req.Name,
		buttonCols.MenuId: req.MenuId,
	}).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, consts.ErrButtonNameExists
	}
	_, err = dao.Button.Ctx(ctx).Data(g.Map{
		buttonCols.MenuId:    req.MenuId,
		buttonCols.MenuTitle: req.MenuTitle,
		buttonCols.Name:      req.Name,
		buttonCols.Title:     req.Title,
		buttonCols.Remark:    req.Remark,
	}).Insert()
	if err != nil {
		return nil, err
	}
	return
}

func UpdateButton(ctx context.Context, req *vbutton.UpdateButtonReq) (res *vbutton.UpdateButtonRes, err error) {
	_, err = dao.Button.Ctx(ctx).Where(buttonCols.Id, req.Id).Data(g.Map{
		buttonCols.MenuId:    req.MenuId,
		buttonCols.MenuTitle: req.MenuTitle,
		buttonCols.Name:      req.Name,
		buttonCols.Title:     req.Title,
		buttonCols.Remark:    req.Remark,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func DeleteButton(ctx context.Context, req *vbutton.DeleteButtonReq) (res *vbutton.DeleteButtonRes, err error) {
	_, err = dao.Button.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	// 删除关联的casbin，删除角色与按钮关联
	var obj = "按钮" + gconv.String(req.Id)
	_, err = xcasbin.Enforcer.RemoveFilteredPolicy(1, obj)
	if err != nil {
		return nil, err
	}
	return
}

func ListButton(ctx context.Context, req *vbutton.ListButtonReq) (res *vbutton.ListButtonRes, err error) {
	var resp = &vbutton.ListButtonRes{
		List:          make([]*entity.Button, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
	var data = g.Map{}
	if req.Name != "" {
		data[buttonCols.Name+" like ?"] = "%" + req.Name + "%"
	}
	if req.Title != "" {
		data[buttonCols.Title+" like ?"] = "%" + req.Title + "%"
	}
	if req.MenuTitle != "" {
		data[buttonCols.MenuTitle+" like ?"] = "%" + req.MenuTitle + "%"
	}
	if req.Id != 0 {
		data[buttonCols.Id] = req.Id
	}
	var model = dao.Button.Ctx(ctx).Where(data).Order(buttonCols.MenuId)
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

func OneButton(ctx context.Context, req *vbutton.OneButtonReq) (res *vbutton.OneButtonRes, err error) {
	err = dao.Button.Ctx(ctx).Where(buttonCols.Id, req.Id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}
