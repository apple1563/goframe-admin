package regionService

import (
	"context"
	"goframe-starter/api/vregion"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
	"goframe-starter/utility/xcache"
)

func GetTreeList(ctx context.Context, req *vregion.ListTreeReq) (res *vregion.ListTreeRes, err error) {
	value, err := xcache.Instance.GetOrSetFunc(ctx, "xcache:region", func(ctx context.Context) (value interface{}, err error) {
		var list []*entity.Provinces
		err = dao.Provinces.Ctx(ctx).Scan(&list)
		if err != nil {
			return nil, err
		}
		res = &vregion.ListTreeRes{
			List: GenTreeList(list, 0),
		}
		return res, nil
	}, 0)
	if err != nil {
		return nil, err
	}
	err = value.Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}

func GenTreeList(list []*entity.Provinces, pid int64) (res []*vregion.TreeItem) {
	res = make([]*vregion.TreeItem, 0)
	for _, item := range list {
		if item.Pid == pid {
			var obj = &vregion.TreeItem{
				Provinces: item,
			}
			obj.Children = GenTreeList(list, item.Id)
			res = append(res, obj)
		}
	}
	return
}
