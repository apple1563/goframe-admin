package menu

import (
	"context"
	"goframe-starter/api/vmenu"
	"goframe-starter/internal/service/menuService"
)

type Menu struct{}

var MenuCtrl = new(Menu)

func (u *Menu) AddMenu(ctx context.Context, req *vmenu.AddMenuReq) (res *vmenu.AddMenuRes, err error) {
	return menuService.AddMenu(ctx, req)
}

func (u *Menu) ListMenu(ctx context.Context, req *vmenu.ListMenuReq) (res *vmenu.ListMenuRes, err error) {
	return menuService.ListMenu(ctx, req)
}
