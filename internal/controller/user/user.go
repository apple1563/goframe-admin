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

func (u *User) ListUser(ctx context.Context, req *vuser.ListUserReq) (res *vuser.ListUserRes, err error) {
	return userService.ListUser(ctx, req)
}

func (u *User) UserInfo(ctx context.Context, req *vuser.UserInfoReq) (res *vuser.UserInfoRes, err error) {
	return userService.UserInfo(ctx, req)
}
