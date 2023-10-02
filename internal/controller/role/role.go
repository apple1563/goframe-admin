package role

import (
	"context"
	"goframe-starter/api/vrole"
	"goframe-starter/internal/service/roleService"
)

type Role struct{}

var Ctrl = new(Role)

func (u *Role) AddRole(ctx context.Context, req *vrole.AddRoleReq) (res *vrole.AddRoleRes, err error) {
	return roleService.AddRole(ctx, req)
}

func (u *Role) DeleteRole(ctx context.Context, req *vrole.DeleteRoleReq) (res *vrole.DeleteRoleRes, err error) {
	return roleService.DeleteRole(ctx, req)
}

func (u *Role) UpdateRole(ctx context.Context, req *vrole.UpdateRoleReq) (res *vrole.UpdateRoleRes, err error) {
	return roleService.UpdateRole(ctx, req)
}

func (u *Role) ListRole(ctx context.Context, req *vrole.ListRoleReq) (res *vrole.ListRoleRes, err error) {
	return roleService.ListRole(ctx, req)
}

func (u *Role) OneRole(ctx context.Context, req *vrole.OneRoleReq) (res *vrole.OneRoleRes, err error) {
	return roleService.OneRole(ctx, req)
}
