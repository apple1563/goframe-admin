package cmd

import (
	"context"
	"goframe-starter/internal/controller/api"
	"goframe-starter/internal/controller/button"
	"goframe-starter/internal/controller/dict"
	"goframe-starter/internal/controller/file"
	"goframe-starter/internal/controller/log"
	"goframe-starter/internal/controller/menu"
	"goframe-starter/internal/controller/otp"
	"goframe-starter/internal/controller/region"
	"goframe-starter/internal/controller/role"
	"goframe-starter/internal/controller/user"
	"goframe-starter/internal/middleware"
	"goframe-starter/utility/xcache"
	"goframe-starter/utility/xcasbin"
	"goframe-starter/utility/xgtoken"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server(g.Cfg().MustGet(ctx, "server.name").String())
			s.BindMiddlewareDefault(middleware.DefaultCORS, middleware.RequestIpLimit, middleware.AutoAddApi)
			xcasbin.Init()
			xcache.Init()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.Ctrl,
						menu.Ctrl,
						role.Ctrl,
						button.Ctrl,
						region.Ctrl,
						dict.Ctrl,
						api.Ctrl,
						log.Ctrl,
						file.Ctrl,
						otp.Ctrl,
					)
				})
			})
			xgtoken.InitGtoken(ctx)
			s.Run()
			return nil
		},
	}
)
