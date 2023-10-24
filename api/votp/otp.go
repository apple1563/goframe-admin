package votp

import "github.com/gogf/gf/v2/frame/g"

type AddOtpReq struct {
	g.Meta `tags:"OTP校验码" method:"post" path:"/otp"  sm:"添加" dc:"添加" `
	Secret string `json:"secret" v:"required"`
	Code   string `json:"code" v:"required" dc:"输入的校验码"`
}
type AddOtpRes struct{}

type UpdateOtpReq struct {
	g.Meta `tags:"OTP校验码" method:"put" path:"/otp"  sm:"更新" dc:"更新" `
	Secret string `json:"secret" v:"required"`
	Code   string `json:"code" v:"required" dc:"输入的校验码"`
}
type UpdateOtpRes struct{}

type UpdateOtpStatusReq struct {
	g.Meta `tags:"OTP校验码" method:"put" path:"/otp/status"  sm:"更新" dc:"更新开启与否" `
	Status int `json:"status" `
}
type UpdateOtpStatusRes struct{}

type OtpReq struct {
	g.Meta `tags:"OTP校验码" method:"get" path:"/otp"  sm:"获取" dc:"获取当前用户secret" `
}
type OtpRes struct {
	Secret string `json:"secret"`
	Url    string `json:"url"`
}

type OtpCheckReq struct {
	g.Meta `tags:"OTP校验码" method:"get" path:"/otp/check"  sm:"获取" dc:"获取当前用户是否开启otp验证" `
}
type OtpCheckRes struct {
	IsOpened bool `json:"isOpened"`
	IsSetted bool `json:"isSetted" dc:"是否设置了otp校验"`
}

type OtpCheckByUsernameReq struct {
	g.Meta   `tags:"OTP校验码" method:"get" path:"/otp/username"  sm:"获取" dc:"通过用户名获取当前用户是否开启otp验证，登录时用" `
	Username string `json:"username" v:"required"    description:""`
}
type OtpCheckByUsernameRes struct {
	IsOpened bool `json:"isOpened"`
}

type OtpValidateReq struct {
	g.Meta   `tags:"OTP校验码" method:"post" path:"/otp/validate"  sm:"校验" dc:"otp验证" `
	Username string `json:"username"    description:""`
	Code     string `json:"code" v:"required" dc:"输入的校验码"`
}
type OtpValidateRes struct {
}
