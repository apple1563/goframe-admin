package dictService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vcommon"
	"goframe-starter/api/vdict"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
)

var dictCols = dao.Dict.Columns()

func AddDict(ctx context.Context, req *vdict.AddDictReq) (res *vdict.AddDictRes, err error) {
	var data = g.Map{}
	data[dictCols.CreateBy] = ctx.Value("uid")
	if req.Remark != "" {
		data[dictCols.Remark] = req.Remark
	}
	if req.ConfigType != "" {
		data[dictCols.ConfigType] = req.ConfigType
	}
	if req.ConfigName != "" {
		data[dictCols.ConfigName] = req.ConfigName
	}
	if req.ConfigKey != "" {
		data[dictCols.ConfigKey] = req.ConfigKey
	}
	if req.ConfigValue != "" {
		data[dictCols.ConfigValue] = req.ConfigValue
	}
	_, err = dao.Dict.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, err
	}
	return
}

func UpdateDict(ctx context.Context, req *vdict.UpdateDictReq) (res *vdict.UpdateDictRes, err error) {
	var data = g.Map{}
	data[dictCols.UpdateBy] = ctx.Value("uid")
	if req.Remark != "" {
		data[dictCols.Remark] = req.Remark
	}
	if req.ConfigType != "" {
		data[dictCols.ConfigType] = req.ConfigType
	}
	if req.ConfigName != "" {
		data[dictCols.ConfigName] = req.ConfigName
	}
	if req.ConfigKey != "" {
		data[dictCols.ConfigKey] = req.ConfigKey
	}
	if req.ConfigValue != "" {
		data[dictCols.ConfigValue] = req.ConfigValue
	}
	_, err = dao.Dict.Ctx(ctx).Where(dictCols.Id, req.Id).Data(data).Update()
	if err != nil {
		return nil, err
	}
	return
}

func DeleteDict(ctx context.Context, req *vdict.DeleteDictReq) (res *vdict.DeleteDictRes, err error) {
	_, err = dao.Dict.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func ListDict(ctx context.Context, req *vdict.ListDictReq) (res *vdict.ListDictRes, err error) {
	var resp = &vdict.ListDictRes{
		List:          make([]*entity.Dict, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
	var data = g.Map{}
	if req.ConfigName != "" {
		data[dictCols.ConfigName+" like ?"] = "%" + req.ConfigName + "%"
	}
	if req.Id != 0 {
		data[dictCols.Id] = req.Id
	}
	if req.ConfigKey != "" {
		data[dictCols.ConfigKey] = req.ConfigKey
	}
	if req.ConfigType != "" {
		data[dictCols.ConfigType] = req.ConfigType
	}
	if req.CreateBy != 0 {
		data[dictCols.CreateBy] = req.CreateBy
	}
	if req.UpdateBy != 0 {
		data[dictCols.UpdateBy] = req.UpdateBy
	}
	var model = dao.Dict.Ctx(ctx).Where(data).Order(dictCols.ConfigType)
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

func OneDict(ctx context.Context, req *vdict.OneDictReq) (res *vdict.OneDictRes, err error) {
	err = dao.Dict.Ctx(ctx).Where(dictCols.Id, req.Id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}

func GetValueByKey(ctx context.Context, key string) string {
	value, err := dao.Dict.Ctx(ctx).Where(dictCols.ConfigKey, key).Fields(dictCols.ConfigValue).Value()
	if err != nil {
		return ""
	}
	return value.String()
}

// 根据上传类型返回字典列表
func ListDictByType(ctx context.Context, configType string) (res *vdict.ListDictByTypeRes, err error) {
	res = &vdict.ListDictByTypeRes{
		List: make([]*entity.Dict, 0),
	}
	err = dao.Dict.Ctx(ctx).Where(dictCols.ConfigType, configType).Scan(&res.List)
	if err != nil {
		return nil, err
	}
	return
}
