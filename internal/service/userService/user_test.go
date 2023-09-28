package userService

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-starter/api/vuser"
	"testing"
)

func TestAdd(t *testing.T) {
	var req = &vuser.AddReq{
		Username: "test",
		Password: "1",
	}
	AddUser(gctx.New(), req)
}
