package dict

import (
	"context"
	"goframe-starter/api/vdict"
	"goframe-starter/internal/service/dictService"
)

type Dict struct{}

var Ctrl = new(Dict)

func (u *Dict) AddDict(ctx context.Context, req *vdict.AddDictReq) (res *vdict.AddDictRes, err error) {
	return dictService.AddDict(ctx, req)
}

func (u *Dict) DeleteDict(ctx context.Context, req *vdict.DeleteDictReq) (res *vdict.DeleteDictRes, err error) {
	return dictService.DeleteDict(ctx, req)
}

func (u *Dict) UpdateDict(ctx context.Context, req *vdict.UpdateDictReq) (res *vdict.UpdateDictRes, err error) {
	return dictService.UpdateDict(ctx, req)
}

func (u *Dict) ListDict(ctx context.Context, req *vdict.ListDictReq) (res *vdict.ListDictRes, err error) {
	return dictService.ListDict(ctx, req)
}

func (u *Dict) OneDict(ctx context.Context, req *vdict.OneDictReq) (res *vdict.OneDictRes, err error) {
	return dictService.OneDict(ctx, req)
}

func (u *Dict) ListDictByType(ctx context.Context, req *vdict.ListDictByTypeReq) (res *vdict.ListDictByTypeRes, err error) {
	return dictService.ListDictByType(ctx, req.ConfigType)
}
