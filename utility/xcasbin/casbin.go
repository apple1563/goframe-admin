package xcasbin

import (
	"context"
	"github.com/dobyte/gf-casbin"
	"github.com/gogf/gf/v2/frame/g"
)

var Enforcer *casbin.Enforcer

func CreateCasbinEnforcer(ctx context.Context) {
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
