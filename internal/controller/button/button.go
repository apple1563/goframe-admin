package button

import (
	"context"
	"goframe-starter/api/vbutton"
	"goframe-starter/internal/service/buttonService"
)

type Button struct{}

var Ctrl = new(Button)

func (u *Button) AddButton(ctx context.Context, req *vbutton.AddButtonReq) (res *vbutton.AddButtonRes, err error) {
	return buttonService.AddButton(ctx, req)
}

func (u *Button) DeleteButton(ctx context.Context, req *vbutton.DeleteButtonReq) (res *vbutton.DeleteButtonRes, err error) {
	return buttonService.DeleteButton(ctx, req)
}

func (u *Button) UpdateButton(ctx context.Context, req *vbutton.UpdateButtonReq) (res *vbutton.UpdateButtonRes, err error) {
	return buttonService.UpdateButton(ctx, req)
}

func (u *Button) ListButton(ctx context.Context, req *vbutton.ListButtonReq) (res *vbutton.ListButtonRes, err error) {
	return buttonService.ListButton(ctx, req)
}

func (u *Button) OneButton(ctx context.Context, req *vbutton.OneButtonReq) (res *vbutton.OneButtonRes, err error) {
	return buttonService.OneButton(ctx, req)
}

func (u *Button) AddButtonForRole(ctx context.Context, req *vbutton.ButtonForRoleReq) (res *vbutton.ButtonForRoleRes, err error) {
	return buttonService.AddButtonForRole(ctx, req)
}

func (u *Button) GetButtonByRole(ctx context.Context, req *vbutton.ButtonByRoleReq) (res *vbutton.ButtonByRoleRes, err error) {
	return buttonService.GetButtonByRole(ctx, req)
}
