package middleware

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yudeguang/ratelimit"
	"time"
)

func MiddlewareDefaultCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func MiddlewareRequestIpLimit(r *ghttp.Request) {
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
