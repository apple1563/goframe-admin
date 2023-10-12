package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ErrTryAgain       = gerror.NewCode(gcode.New(-4, "失败请重试", ""))
	ErrRepeatedSubmit = gerror.NewCode(gcode.New(-5, "请不要重复提交", ""))
	ErrSql            = gerror.NewCode(gcode.New(-6, "sql执行异常", ""))
	ErrTypeAssertion  = gerror.NewCode(gcode.New(-7, "类型断言错误", ""))
	ErrCaptcha        = gerror.NewCode(gcode.New(-8, "验证码错误", ""))

	ErrLogin            = gerror.NewCode(gcode.New(-100, "用户名或密码错误", ""))
	ErrPassEmpty        = gerror.NewCode(gcode.New(-101, "密码不能为空", ""))
	ErrFormatEmail      = gerror.NewCode(gcode.New(-102, "邮箱格式不正确", ""))
	ErrUnameExist       = gerror.NewCode(gcode.New(-103, "用户名已存在", ""))
	ErrUnameFormat      = gerror.NewCode(gcode.New(-104, "用户名长度在4到12位之间", ""))
	ErrPassFormat       = gerror.NewCode(gcode.New(-105, "密码格式为任意可见字符，长度在6~18之间", ""))
	ErrPassErrorTooMany = gerror.NewCode(gcode.New(-106, "密码错误次数太多", ""))
	ErrOldPassNotMatch  = gerror.NewCode(gcode.New(-107, "旧密码不正确", ""))
	ErrNicknameEmpty    = gerror.NewCode(gcode.New(-108, "昵称不能为空", ""))
	ErrMaxLengthSixTy   = gerror.NewCode(gcode.New(-109, "允许的最长字符为16", ""))
	ErrIconEmpty        = gerror.NewCode(gcode.New(-110, "图片不能为空", ""))
	ErrUserDoesNotExist = gerror.NewCode(gcode.New(-111, "用户不存在", ""))
	ErrBalance          = gerror.NewCode(gcode.New(-112, "用户余额错误", ""))
	ErrUserDead         = gerror.NewCode(gcode.New(-113, "账号已注销", ""))
	ErrUserDisable      = gerror.NewCode(gcode.New(-114, "账号已禁用", ""))

	ErrAreaCode       = gerror.NewCode(gcode.New(-216, "手机区号错误", ""))
	ErrPhoneEmpty     = gerror.NewCode(gcode.New(-120, "手机号不能为空", ""))
	ErrPhoneLength5   = gerror.NewCode(gcode.New(-121, "手机号长度最小5位数", ""))
	ErrUsernameExists = gerror.NewCode(gcode.New(-122, "用户名已存在", ""))

	ErrDepositClosed    = gerror.NewCode(gcode.New(-1000, "充值通道已关闭", ""))
	ErrDepositIncorrect = gerror.NewCode(gcode.New(-1001, "充值金额不正确", ""))
	ErrDepositMin       = gerror.NewCode(gcode.New(-1002, "最低充值额度100USDT", ""))

	ErrWithdrawClose       = gerror.NewCode(gcode.New(-2000, "提现通道已关闭", ""))
	ErrWithdrawIncorrect   = gerror.NewCode(gcode.New(-2001, "提现金额不正确", ""))
	ErrWithdrawBindAccount = gerror.NewCode(gcode.New(-2002, "请先绑定提现账号", ""))
	ErrWithdrawMin         = gerror.NewCode(gcode.New(-2003, "低于最低提现额度", ""))
	ErrBindBankcard        = gerror.NewCode(gcode.New(-2004, "您已绑定过银行卡", ""))

	ErrMenuPathExists         = gerror.NewCode(gcode.New(-3000, "该path已存在", ""))
	ErrMenuPathDeleteChildren = gerror.NewCode(gcode.New(-3001, "请先删除该菜单下的所有菜单及按钮", ""))
	ErrButtonNameExists       = gerror.NewCode(gcode.New(-3002, "该按钮标识码已存在", ""))

	ErrRoleNameExists = gerror.NewCode(gcode.New(-4000, "角色名称已存在", ""))
	ErrRoleDelete     = gerror.NewCode(gcode.New(-4001, "角色已被用户关联，请将关联用户设为其他角色", ""))
)

const Role_User_Code = 1
const Role_Agent_Code = 3
const Role_Admin_Code = 2

const USER_STATUS_ENABLE = 1
const USER_STATUS_DISABLE = 2
const USER_STATUS_DEAD = 3

const Cache_Mode_Memory = 1
const Cache_Mode_Redis = 2
