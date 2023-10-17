package buttonService

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
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

func Test_aaa(t *testing.T) {
	var (
		haystack = `/api/sdf/nn`
		needle   = `/api/`
		result1  = gstr.PosI(haystack, needle)
	)
	fmt.Println(result1)
}
