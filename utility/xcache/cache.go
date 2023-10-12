package xcache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-starter/internal/consts"
)

var Instance *gcache.Cache

func Init() {
	var (
		ctx   = gctx.New()
		cache = gcache.New()
	)
	//prefix := g.Cfg().MustGet(ctx, "cache.prefix").Int()
	model := g.Cfg().MustGet(ctx, "cache.model").Int()
	if model == consts.Cache_Mode_Redis {
		// Create redis cache adapter and set it to cache object.
		cache.SetAdapter(gcache.NewAdapterRedis(g.Redis()))
	}
	Instance = cache
}
