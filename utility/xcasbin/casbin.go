package xcasbin

import (
	"github.com/dobyte/gf-casbin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var Enforcer *casbin.Enforcer

func Init() {
	ctx := gctx.New()
	Enf, err := casbin.NewEnforcer(&casbin.Options{
		Model:    g.Cfg().MustGet(ctx, "casbin.modelFile").String(),
		Debug:    false,
		Enable:   true,
		AutoLoad: true,
		Table:    "casbin_policy",
		Link:     g.Cfg().MustGet(ctx, "database.default.link").String(),
	})

	if err != nil {
		panic(err)
	}
	Enforcer = Enf
}
