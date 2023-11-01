package apiService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-starter/api/vapi"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
	"goframe-starter/utility/xcasbin"
)

var apiCols = dao.Api.Columns()

func AddApi(ctx context.Context, req *vapi.AddApiReq) (res *vapi.AddApiRes, err error) {
	_, err = dao.Api.Ctx(ctx).Data(g.Map{
		apiCols.Url:    req.Url,
		apiCols.Group:  req.Group,
		apiCols.Method: req.Method,
		apiCols.Remark: req.Remark,
	}).Insert()
	if err != nil {
		return nil, err
	}
	return
}

func UpdateApi(ctx context.Context, req *vapi.UpdateApiReq) (res *vapi.UpdateApiRes, err error) {
	_, err = dao.Api.Ctx(ctx).Where(apiCols.Id, req.Id).Data(g.Map{
		apiCols.Url:    req.Url,
		apiCols.Group:  req.Group,
		apiCols.Method: req.Method,
		apiCols.Remark: req.Remark,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func DeleteApi(ctx context.Context, req *vapi.DeleteApiReq) (res *vapi.DeleteApiRes, err error) {
	var api *entity.Api
	err = dao.Api.Ctx(ctx).Where(apiCols.Id, req.Id).Scan(&api)
	if err != nil {
		return nil, err
	}
	_, err = dao.Api.Ctx(ctx).Where(apiCols.Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}
	// 删除关联的casbin，删除角色与按钮关联
	var obj = api.Url
	var act = api.Method
	_, err = xcasbin.Enforcer.RemoveFilteredPolicy(1, obj, act)
	if err != nil {
		return nil, err
	}
	return
}

func ListApi(ctx context.Context, req *vapi.ListApiReq) (res *vapi.ListApiRes, err error) {
	var resp = &vapi.ListApiRes{
		List:          make([]*entity.Api, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
	var data = g.Map{}
	if req.Url != "" {
		data[apiCols.Url+" like ?"] = "%" + req.Url + "%"
	}
	if req.Group != "" {
		data[apiCols.Group] = req.Group
	}
	var model = dao.Api.Ctx(ctx).Where(data).OrderDesc(apiCols.CreatedAt).OrderAsc(apiCols.Group)
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

func OneApi(ctx context.Context, req *vapi.OneApiReq) (res *vapi.OneApiRes, err error) {
	err = dao.Api.Ctx(ctx).Where(apiCols.Id, req.Id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}

func AddApiForRole(ctx context.Context, req *vapi.ApiForRoleReq) (res *vapi.ApiForRoleRes, err error) {
	var sub = consts.Role_Api_Prefix + gconv.String(req.RoleId)
	_, err = xcasbin.Enforcer.RemoveFilteredPolicy(0, sub)
	if err != nil {
		return nil, err
	}
	for _, api := range req.Apis {
		_, err := xcasbin.Enforcer.AddPolicy(sub, api.Url, api.Method)
		if err != nil {
			return nil, err
		}
	}
	return
}
func GetApiByRole(ctx context.Context, req *vapi.ApiByRoleReq) (res *vapi.ApiByRoleRes, err error) {
	var sub = consts.Role_Api_Prefix + gconv.String(req.RoleId)
	var resp = &vapi.ApiByRoleRes{
		List: make([]*vapi.ApiByRole, 0),
	}
	var rules = xcasbin.Enforcer.GetFilteredPolicy(0, sub)
	for _, rule := range rules {
		var apiByRole = &vapi.ApiByRole{
			Url:    rule[1],
			Method: rule[2],
		}
		resp.List = append(resp.List, apiByRole)
	}
	return resp, nil
}

func CheckApiExists(ctx context.Context, url string, method string) bool {
	count, err := dao.Api.Ctx(ctx).Where(g.Map{
		apiCols.Url:    url,
		apiCols.Method: method,
	}).Count()
	if err != nil {
		return true
	}
	if count > 0 {
		return true
	}
	return false
}
