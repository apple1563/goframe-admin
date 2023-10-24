package otp

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-starter/api/votp"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/do"
	"goframe-starter/internal/model/entity"
	"goframe-starter/utility/xotp"
)

type MyOtp struct{}

var Ctrl = new(MyOtp)
var cols = dao.OtpAuth.Columns()

func (u *MyOtp) AddOtp(ctx context.Context, req *votp.AddOtpReq) (res *votp.AddOtpRes, err error) {
	var flag = xotp.ValidateOtp(req.Code, req.Secret)
	if flag {
		_, err = dao.OtpAuth.Ctx(ctx).Data(do.OtpAuth{
			Uid:      ctx.Value("uid"),
			Username: ctx.Value("username"),
			Secret:   req.Secret,
		}).Insert()
		if err != nil {
			return nil, err
		}
		return
	} else {
		return nil, consts.ErrOtpCode
	}
}

func (u *MyOtp) UpdateOtp(ctx context.Context, req *votp.UpdateOtpReq) (res *votp.UpdateOtpRes, err error) {
	var flag = xotp.ValidateOtp(req.Code, req.Secret)
	if flag {
		var data = g.Map{}
		data[cols.Secret] = req.Secret
		data[cols.Uid] = gconv.Uint(ctx.Value("uid"))
		_, err = dao.OtpAuth.Ctx(ctx).Where(cols.Uid, ctx.Value("uid")).Data(data).Update()
		if err != nil {
			return nil, err
		}
		return
	} else {
		return nil, consts.ErrOtpCode
	}
}

func (u *MyOtp) UpdateOtpStatus(ctx context.Context, req *votp.UpdateOtpStatusReq) (res *votp.UpdateOtpStatusRes, err error) {
	_, err = dao.OtpAuth.Ctx(ctx).Where(cols.Uid, ctx.Value("uid")).Data(g.Map{
		cols.Status: req.Status,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func (u *MyOtp) GetOtp(ctx context.Context, req *votp.OtpReq) (res *votp.OtpRes, err error) {
	res = &votp.OtpRes{}
	/*var count int
	err = dao.OtpAuth.Ctx(ctx).Where(cols.Uid, ctx.Value("uid")).ScanAndCount(&res, &count, false)
	if err != nil {
		return nil, err
	}
	// 没有设置过，则生成
	if count < 1 {*/
	otp, err := xotp.GenerateOtp(gconv.String(ctx.Value("uid")), "goframe-admin")
	if err != nil {
		return nil, err
	}
	res.Url = otp.String()
	res.Secret = otp.Secret()
	//}

	return
}

func (u *MyOtp) OtpCheck(ctx context.Context, req *votp.OtpCheckReq) (res *votp.OtpCheckRes, err error) {
	res = &votp.OtpCheckRes{}
	var data *entity.OtpAuth
	var count int
	err = dao.OtpAuth.Ctx(ctx).Where(cols.Uid, ctx.Value("uid")).ScanAndCount(&data, &count, false)
	if err != nil {
		return nil, err
	}
	if count < 1 {
		res.IsSetted = false
		res.IsOpened = false
	} else {
		res.IsSetted = true
		res.IsOpened = data.Status == consts.OTP_STATUS_ON
	}
	return
}

func (u *MyOtp) OtpCheckByUsername(ctx context.Context, req *votp.OtpCheckByUsernameReq) (res *votp.OtpCheckByUsernameRes, err error) {
	res = &votp.OtpCheckByUsernameRes{}
	var data *entity.OtpAuth
	var count int
	err = dao.OtpAuth.Ctx(ctx).Where(cols.Username, req.Username).ScanAndCount(&data, &count, false)
	if err != nil {
		return nil, err
	}
	if count < 1 {
		res.IsOpened = false
	} else {
		res.IsOpened = data.Status == consts.OTP_STATUS_ON
	}
	return
}

func (u *MyOtp) OtpValidate(ctx context.Context, req *votp.OtpValidateReq) (res *votp.OtpValidateRes, err error) {
	res = &votp.OtpValidateRes{}
	var data *entity.OtpAuth
	err = dao.OtpAuth.Ctx(ctx).Where(cols.Username, req.Username).Scan(&data)
	if err != nil {
		return nil, err
	}
	var flag = xotp.ValidateOtp(req.Code, data.Secret)
	if flag {
		return
	} else {
		return nil, consts.ErrOtpCode
	}
}
