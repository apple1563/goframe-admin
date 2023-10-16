package api

import (
	"context"
	"goframe-starter/api/vapi"
	"goframe-starter/internal/service/apiService"
)

type Api struct{}

var Ctrl = new(Api)

func (u *Api) AddApi(ctx context.Context, req *vapi.AddApiReq) (res *vapi.AddApiRes, err error) {
	return apiService.AddApi(ctx, req)
}

func (u *Api) DeleteApi(ctx context.Context, req *vapi.DeleteApiReq) (res *vapi.DeleteApiRes, err error) {
	return apiService.DeleteApi(ctx, req)
}

func (u *Api) UpdateApi(ctx context.Context, req *vapi.UpdateApiReq) (res *vapi.UpdateApiRes, err error) {
	return apiService.UpdateApi(ctx, req)
}

func (u *Api) ListApi(ctx context.Context, req *vapi.ListApiReq) (res *vapi.ListApiRes, err error) {
	return apiService.ListApi(ctx, req)
}

func (u *Api) OneApi(ctx context.Context, req *vapi.OneApiReq) (res *vapi.OneApiRes, err error) {
	return apiService.OneApi(ctx, req)
}

func (u *Api) AddApiForRole(ctx context.Context, req *vapi.ApiForRoleReq) (res *vapi.ApiForRoleRes, err error) {
	return apiService.AddApiForRole(ctx, req)
}

func (u *Api) GetApiByRole(ctx context.Context, req *vapi.ApiByRoleReq) (res *vapi.ApiByRoleRes, err error) {
	return apiService.GetApiByRole(ctx, req)
}
