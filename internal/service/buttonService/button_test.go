package buttonService

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"goframe-starter/internal/dao"
	"testing"
	"time"
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
	var ctx = gctx.New()
	var folder = time.Now().Format("2006-01-02")
	var dirPath = g.Cfg().MustGet(ctx, "server.uploadRoot").String() + "/" + folder
	err := gfile.Remove(gfile.MainPkgPath() + "/" + dirPath + "/" + "cwbcxlyrbijkzaspxn.sql")
	if err != nil {
		g.Dump(err)
		return
	}
	var (
		haystack = `/api/sdf/nn`
		needle   = `/api/`
		result1  = gstr.PosI(haystack, needle)
	)
	fmt.Println(result1)
}

func Test_b(t *testing.T) {
	//  leftjoin跟innerjoin返回的结果是一样的
	array, err := g.Model("user u").LeftJoin("user_relation ud", "u.id=ud.user_id").Fields("u.id").Where("ud.user_id", 8).Array()
	if err != nil {
		return
	}
	g.Dump(array)
	array2, err := g.Model("user u").InnerJoin("user_relation ud", "u.id=ud.user_id").Fields("u.id").Where("ud.user_id", 8).Array()
	if err != nil {
		return
	}
	g.Dump(array2)
}
func Test111(t *testing.T) {
	g.Dump(gfile.Ext("go1.21.2.linux-amd64.tar.gz"))
}
