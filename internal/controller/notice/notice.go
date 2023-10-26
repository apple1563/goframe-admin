package notice

import (
	"context"
	"goframe-starter/api/vnotice"
	"goframe-starter/internal/service/noticeService"
)

type Notice struct{}

var Ctrl = new(Notice)

func (u *Notice) AddNotice(ctx context.Context, req *vnotice.AddNoticeReq) (res *vnotice.AddNoticeRes, err error) {
	return noticeService.AddNotice(ctx, req)
}

func (u *Notice) DeleteNotice(ctx context.Context, req *vnotice.DeleteNoticeReq) (res *vnotice.DeleteNoticeRes, err error) {
	return noticeService.DeleteNotice(ctx, req)
}

func (u *Notice) UpdateNotice(ctx context.Context, req *vnotice.UpdateNoticeReq) (res *vnotice.UpdateNoticeRes, err error) {
	return noticeService.UpdateNotice(ctx, req)
}

func (u *Notice) ListNoticeForSender(ctx context.Context, req *vnotice.ListNoticeReq) (res *vnotice.ListNoticeRes, err error) {
	return noticeService.ListNoticeForSender(ctx, req)
}

func (u *Notice) OneNotice(ctx context.Context, req *vnotice.OneNoticeReq) (res *vnotice.OneNoticeRes, err error) {
	return noticeService.OneNotice(ctx, req)
}

func (u *Notice) ListNoticeForReceiver(ctx context.Context, req *vnotice.ListNoticeForReceiverReq) (res *vnotice.ListNoticeForReceiverRes, err error) {
	return noticeService.ListNoticeForReceiver(ctx, req)
}
