package noticeService

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-starter/api/vcommon"
	"goframe-starter/api/vnotice"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/do"
	"goframe-starter/internal/model/entity"
	"sync"
)

var noticeCols = dao.Notice.Columns()

func AddNotice(ctx context.Context, req *vnotice.AddNoticeReq) (res *vnotice.AddNoticeRes, err error) {
	g.DumpWithType(req.Receivers)
	if req.Receivers == "" {
		return nil, consts.ErrNoticeReceivers
	}
	if req.Content == "" {
		return nil, consts.ErrNoticeContent
	}
	var data = g.Map{}
	data[noticeCols.Creater] = ctx.Value("username")
	data[noticeCols.Receivers] = req.Receivers
	data[noticeCols.Content] = req.Content
	if req.Title != "" {
		data[noticeCols.Title] = req.Title
	}
	if req.Tag != "" {
		data[noticeCols.Tag] = req.Tag
	}
	if req.Sort != 0 {
		data[noticeCols.Sort] = req.Sort
	}
	if req.Remark != "" {
		data[noticeCols.Remark] = req.Remark
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		noticeId, err := dao.Notice.Ctx(ctx).TX(tx).Data(data).InsertAndGetId()
		if err != nil {
			return err
		}
		var receivers []uint
		err = json.Unmarshal([]byte(req.Receivers), &receivers)
		if err != nil {
			return err
		}
		//  插入notice_relation表  用事务
		var wg sync.WaitGroup
		var errChan = make(chan error)
		for _, receiver := range receivers {
			wg.Add(1)
			go func(receiver uint) {
				defer wg.Done()
				_, err := dao.NoticeUserRelation.Ctx(ctx).TX(tx).Data(do.NoticeUserRelation{
					NoticeId: noticeId,
					Uid:      receiver,
				}).Insert()
				if err != nil {
					errChan <- err
				}
			}(receiver)
		}
		go func() {
			wg.Wait()
			close(errChan)
		}()
		for ec := range errChan {
			//todo 可以考虑将发送失败的用户放到定时任务里，直到发送成功
			g.Log().Error(ctx, "发送通告失败", ec)
			//return ec
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return
}

func UpdateNotice(ctx context.Context, req *vnotice.UpdateNoticeReq) (res *vnotice.UpdateNoticeRes, err error) {
	var data = g.Map{}
	if req.Receivers != "" {
		data[noticeCols.Receivers] = req.Receivers
	}
	if req.Title != "" {
		data[noticeCols.Title] = req.Title
	}
	if req.Content != "" {
		data[noticeCols.Content] = req.Content
	}
	if req.Tag != "" {
		data[noticeCols.Tag] = req.Tag
	}
	if req.Sort != 0 {
		data[noticeCols.Sort] = req.Sort
	}
	if req.Remark != "" {
		data[noticeCols.Remark] = req.Remark
	}
	_, err = dao.Notice.Ctx(ctx).Where(noticeCols.Id, req.Id).Data(data).Update()
	if err != nil {
		return nil, err
	}
	return
}

func DeleteNotice(ctx context.Context, req *vnotice.DeleteNoticeReq) (res *vnotice.DeleteNoticeRes, err error) {
	_, err = dao.Notice.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func ListNoticeForSender(ctx context.Context, req *vnotice.ListNoticeReq) (res *vnotice.ListNoticeRes, err error) {
	var resp = &vnotice.ListNoticeRes{
		List:          make([]*entity.Notice, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
	var data = g.Map{}
	var roleId = gconv.Int(ctx.Value("roleId"))
	if roleId == consts.Role_Root_Code {
		// 超级管理员可以看所有的公告，非超管只能看自己创建的公告
	} else {
		data[noticeCols.Creater] = ctx.Value("username")
	}
	if req.Title != "" {
		data[noticeCols.Title+" like ?"] = "%" + req.Title + "%"
	}
	if req.Content != "" {
		data[noticeCols.Content+" like ?"] = "%" + req.Content + "%"
	}
	if req.Tag != "" {
		data[noticeCols.Tag+" like ?"] = "%" + req.Tag + "%"
	}
	if req.Creater != "" {
		data[noticeCols.Creater+" like ?"] = "%" + req.Creater + "%"
	}
	if req.Id != 0 {
		data[noticeCols.Id] = req.Id
	}

	var model = dao.Notice.Ctx(ctx).Where(data).OrderAsc(noticeCols.Sort).OrderDesc(noticeCols.CreatedAt)
	if req.PageSize != 0 {
		resp.PageIndex = req.PageIndex
		resp.PageSize = req.PageSize
		model = model.Page(req.PageIndex, req.PageSize)
	}
	err = model.ScanAndCount(&resp.List, &resp.Total, false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func OneNotice(ctx context.Context, req *vnotice.OneNoticeReq) (res *vnotice.OneNoticeRes, err error) {
	err = dao.Notice.Ctx(ctx).Where(noticeCols.Id, req.Id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}

// 接收端
var noticeForReceiverCols = dao.NoticeUserRelation.Columns()

func ListNoticeForReceiver(ctx context.Context, req *vnotice.ListNoticeForReceiverReq) (res *vnotice.ListNoticeForReceiverRes, err error) {
	//  在添加公告的时候在公告与个人的关联表写入数据，查询的时候join这个表
	var resp = &vnotice.ListNoticeForReceiverRes{
		List:          make([]*vnotice.ItemNoticeForReceiver, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
	var uid = ctx.Value("uid")
	var model = g.Model(dao.NoticeUserRelation.Table(), "a").InnerJoin(dao.Notice.Table(), "b", "a.notice_id=b.id").Fields("a.*,b.remark,b.title,b.tag,b.content,b.sort").Where(
		"a.uid", uid)
	if req.Title != "" {
		model = model.WhereLike("b.title", req.Title)
	}
	if req.Content != "" {
		model = model.WhereLike("b.content", req.Content)
	}
	if req.Tag != "" {
		model = model.WhereLike("b.tag", req.Tag)
	}
	if req.Status != 0 {
		model = model.Where("a.status", req.Status)
	}

	model = model.OrderAsc("b.sort").OrderDesc("b.created_at").OrderAsc("a.status")
	if req.PageSize != 0 {
		resp.PageIndex = req.PageIndex
		resp.PageSize = req.PageSize
		model = model.Page(req.PageIndex, req.PageSize)
	}
	err = model.ScanAndCount(&resp.List, &resp.Total, false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DeleteNoticeForReceiver(ctx context.Context, req *vnotice.DeleteNoticeForReceiverReq) (res *vnotice.DeleteNoticeForReceiverRes, err error) {
	_, err = dao.NoticeUserRelation.Ctx(ctx).Where(noticeForReceiverCols.Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func UpdateNoticeStatusForReceiver(ctx context.Context, req *vnotice.UpdateNoticeForReceiverReq) (res *vnotice.UpdateNoticeForReceiverRes, err error) {
	_, err = dao.NoticeUserRelation.Ctx(ctx).Where(noticeForReceiverCols.Id, req.Id).Data(g.Map{
		noticeForReceiverCols.Status: req.Status,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func GetNoticeUnreadCountForReceiver(ctx context.Context, req *vnotice.GetNoticeUnreadCountForReceiverReq) (res *vnotice.GetNoticeUnreadCountForReceiverRes, err error) {
	res = &vnotice.GetNoticeUnreadCountForReceiverRes{
		Count: 0,
	}
	count, err := dao.NoticeUserRelation.Ctx(ctx).Where(noticeForReceiverCols.Uid, ctx.Value("uid")).Count(noticeForReceiverCols.Status, 1)
	if err != nil {
		return nil, err
	}
	res.Count = count
	return
}
