package buttonService

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-starter/internal/dao"
	"testing"
)

func TestGroup(t *testing.T) {
	var ctx = gctx.New()
	value, err := dao.Button.Ctx(ctx).Fields("COUNT(*) total,menu_id").Group("menu_id").All()
	if err != nil {
		return
	}
	g.Dump(value)
}
