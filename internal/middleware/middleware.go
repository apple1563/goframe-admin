package middleware

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/yudeguang/ratelimit"
	"goframe-starter/api/vapi"
	"goframe-starter/internal/model/entity"
	"goframe-starter/internal/service/apiService"
	"strings"
	"time"
)

func DefaultCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func RequestIpLimit(r *ghttp.Request) {
	ip := r.GetClientIp()
	rule := ratelimit.NewRule()
	rule.AddRule(time.Hour, 10000)
	rule.AddRule(time.Minute, 600)
	rule.AddRule(time.Second, 10)
	ok := rule.AllowVisit(ip)
	if !ok {
		g2 := gtoken.Resp{
			Code: -999,
			Msg:  "YOUR ACCESS IS ABNORMAL",
		}
		r.Response.WriteJsonExit(g2)

	}
	r.Middleware.Next()
}

func AutoAddApi(r *ghttp.Request) {
	r.Middleware.Next()
	var ctx = gctx.New()
	flag := apiService.CheckApiExists(ctx, r.URL.Path, r.Method)
	if flag {
		return
	}
	//  自动添加api
	var req = &vapi.AddApiReq{
		Api: &entity.Api{
			Url:    r.URL.Path,
			Method: r.Method,
			Group:  strings.Join(strings.Split(r.URL.Path, "/")[:3], "/"),
		},
	}
	_, _ = apiService.AddApi(ctx, req)
}
