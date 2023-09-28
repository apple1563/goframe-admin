package gtokenService

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
	"goframe-starter/utility/xcasbin"
	"goframe-starter/utility/xpwd"
)

type GFtokenFn struct {
}

var GFtokenFnInstance = new(GFtokenFn)

func (*GFtokenFn) LoginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	var (
		ctx      = r.Context()
		username = r.Get("username").String()
		password = r.Get("password").String()
	)
	if username == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT OR PASSWORD CANNOT BE EMPTY."))
		r.ExitAll()
	}
	var user *entity.User
	var count int
	var userCols = dao.User.Columns()
	//  检验账号密码
	err := dao.User.Ctx(ctx).Where(g.Map{
		userCols.Username: username,
	}).ScanAndCount(&user, &count, false)
	if err != nil {
		r.Response.WriteJson(err)
		r.ExitAll()
	}
	if count < 1 {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT ERROR."))
		r.ExitAll()
	}
	if !xpwd.ComparePassword(user.Password, password) {
		r.Response.WriteJson(gtoken.Fail("WRONG PASSWORD."))
		r.ExitAll()
	}
	//  判断角色有无登录后台权限
	enforce, _ := xcasbin.Enforcer.Enforce(user.RoleId, r.URL.Path, r.Method)
	if !enforce {
		r.Response.WriteJson(gtoken.Fail("ACCOUNT ROLE ERROR."))
		r.ExitAll()
	}
	// 写入登录日志
	var loginLogCols = dao.LoginLog.Columns()
	_, err = dao.LoginLog.Ctx(ctx).Data(g.Map{
		loginLogCols.Uid:         user.Id,
		loginLogCols.Role:        user.RoleId,
		loginLogCols.Username:    user.Username,
		loginLogCols.ClientAgent: r.UserAgent(),
		loginLogCols.PRole:       user.PRoleId,
		loginLogCols.Pid:         user.Pid,
		loginLogCols.PUsername:   user.PUsername,
		loginLogCols.Ip:          r.GetClientIp(),
	}).Insert()
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("SAVE LOGIN LOG  ERROR."))
		r.ExitAll()
	}

	return username, user
}

/*
	func AuthBeforeFunc(r *ghttp.Request) bool {
		r.SetCtxVar("time", time.Now())
		return true
	}
*/
func (*GFtokenFn) AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Code != 0 {
		switch r.Method {
		case "PUT", "DELETE", "GET", "POST":
			r.Response.WriteJsonExit(respData)
		}
		return
	}
	var userInfo *entity.User
	err := gjson.New(respData.Data).Get("data").Struct(&userInfo)
	if err != nil {
		r.Response.WriteJsonExit(err)
		return
	}
	r.SetCtxVar("userInfo", userInfo)
	//  登录校验后  casbin权限校验
	//enforce, _ := xcasbin.Cb.Enforce(username, path, method)
	enforce, _ := xcasbin.Enforcer.Enforce(userInfo.RoleId, r.URL.Path, r.Method)
	if !enforce {
		respData.Code = -403
		respData.Data = "未授权"
		r.Response.Status = 403
		r.Response.WriteJsonExit(respData)
		return
	}

	r.Middleware.Next()
	//elapsedTime := time.Since(timeStart)
	//requestBody := r.GetBodyString()
	//写入操作日志
	/*go func(ctx1 context.Context) {
		log := model.OperateLog{
			Ip:          u.ClientIp,
			Path:        r.URL.Path,
			Method:      r.Method,
			Account:     u.Account,
			RoleName:    u.RuleName,
			Request:     fmt.Sprint(requestBody),
			Response:    r.Response.BufferString(),
			CreatedAt:   time.Now(),
			ElapsedTime: elapsedTime.Milliseconds(),
		}
		format := time.Now().Format("2006-01-02")
		sprint := fmt.Sprint("admin_operate_log_", format)
		xelastic.Create(ctx1, sprint, log)
	}(context.Background())*/
}
