package user

import (
	"context"
	"goframe-starter/api/vuser"
	"goframe-starter/internal/service/userService"
)

type User struct{}

var Ctrl = new(User)

func (u *User) AddUser(ctx context.Context, req *vuser.AddUserReq) (res *vuser.AddUserRes, err error) {
	return userService.AddUser(ctx, req)
}

func (u *User) DeleteUser(ctx context.Context, req *vuser.DeleteUserReq) (res *vuser.DeleteUserRes, err error) {
	return userService.DeleteUser(ctx, req)
}

func (u *User) UpdateUser(ctx context.Context, req *vuser.UpdateUserReq) (res *vuser.UpdateUserRes, err error) {
	return userService.UpdateUser(ctx, req)
}

func (u *User) ListUser(ctx context.Context, req *vuser.ListUserReq) (res *vuser.ListUserRes, err error) {
	return userService.ListUser(ctx, req)
}

func (u *User) OneUser(ctx context.Context, req *vuser.OneUserReq) (res *vuser.OneUserRes, err error) {
	return userService.OneUser(ctx, req)
}

func (u *User) UserInfo(ctx context.Context, req *vuser.UserInfoReq) (res *vuser.UserInfoRes, err error) {
	return userService.UserInfo(ctx, req)
}

func (u *User) TreeListUserScope(ctx context.Context, req *vuser.TreeListUserReq) (res *vuser.TreeListUserRes, err error) {
	return userService.TreeListUserScope(ctx, req)
}
