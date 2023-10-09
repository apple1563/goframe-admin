package xgtoken

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-starter/internal/service/gtokenService"
)

type GFToken struct{}

var Gtoken = new(GFToken)
var MyGtoken *gtoken.GfToken

type GFTokenFunc interface {
	LoginBeforeFunc(r *ghttp.Request) (string, interface{})
	//LoginAfterFunc(r *ghttp.Request, respData gtoken.Resp)
	//LogoutBeforeFunc(r *ghttp.Request) bool
	//LogoutAfterFunc(r *ghttp.Request, respData gtoken.Resp)
	//AuthBeforeFunc(r *ghttp.Request) bool
	AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp)
}

func (s *GFToken) New(ctx context.Context, fn GFTokenFunc) *gtoken.GfToken {
	gfToken := &gtoken.GfToken{
		ServerName:       g.Cfg().MustGet(ctx, "server.name").String(),
		CacheKey:         g.Cfg().MustGet(ctx, "gfToken.cacheKey").String(),
		CacheMode:        g.Cfg().MustGet(ctx, "gfToken.cacheMode").Int8(),
		Timeout:          g.Cfg().MustGet(ctx, "gfToken.timeOut").Int(),
		EncryptKey:       g.Cfg().MustGet(ctx, "gfToken.encryptKey").Bytes(),
		MaxRefresh:       g.Cfg().MustGet(ctx, "gfToken.maxRefresh").Int(),
		MultiLogin:       g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool(),
		GlobalMiddleware: g.Cfg().MustGet(ctx, "gfToken.globalMiddleware").Bool(),
		LoginPath:        g.Cfg().MustGet(ctx, "gfToken.loginPath").String(),
		LogoutPath:       g.Cfg().MustGet(ctx, "gfToken.logoutPath").String(),
		AuthPaths:        g.Cfg().MustGet(ctx, "gfToken.authPaths").Strings(),
		AuthExcludePaths: g.Cfg().MustGet(ctx, "gfToken.excludePaths").Strings(),
		LoginBeforeFunc:  fn.LoginBeforeFunc,
		AuthAfterFunc:    fn.AuthAfterFunc,
	}
	g.Dump(gfToken)
	return gfToken
}

func InitGtoken(ctx context.Context) {
	MyGtoken = Gtoken.New(ctx, gtokenService.GFtokenFnInstance)
	err := MyGtoken.Start()
	if err != nil {
		panic(err)
	}
}
