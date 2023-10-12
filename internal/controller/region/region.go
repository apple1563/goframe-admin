package region

import (
	"context"
	"goframe-starter/api/vregion"
	"goframe-starter/internal/service/regionService"
)

type Region struct{}

var Ctrl = new(Region)

func (u *Region) AddRegion(ctx context.Context, req *vregion.ListTreeReq) (res *vregion.ListTreeRes, err error) {
	return regionService.GetTreeList(ctx, req)
}
