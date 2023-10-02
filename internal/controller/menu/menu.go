package menu

import (
	"context"
	"goframe-starter/api/vmenu"
	"goframe-starter/internal/service/menuService"
)

type Menu struct{}

var Ctrl = new(Menu)

func (u *Menu) AddMenu(ctx context.Context, req *vmenu.AddMenuReq) (res *vmenu.AddMenuRes, err error) {
	return menuService.AddMenu(ctx, req)
}

func (u *Menu) DeleteMenu(ctx context.Context, req *vmenu.DeleteMenuReq) (res *vmenu.DeleteMenuRes, err error) {
	return menuService.DeleteMenu(ctx, req)
}

func (u *Menu) UpdateMenu(ctx context.Context, req *vmenu.UpdateMenuReq) (res *vmenu.UpdateMenuRes, err error) {
	return menuService.UpdateMenu(ctx, req)
}

func (u *Menu) ListTreeMenus(ctx context.Context, req *vmenu.ListTReeMenuReq) (res *vmenu.ListTReeMenuRes, err error) {
	return menuService.ListTreeMenus(ctx, req)
}

func (u *Menu) ListVueMenus(ctx context.Context, req *vmenu.VueMenuReq) (res *vmenu.VueMenuRes, err error) {
	return menuService.ListVueMenus(ctx, req)
}
