package xpwd

import "golang.org/x/crypto/bcrypt"

func GenPwd(pwd string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(password)
}

// ComparePassword pwd 数据库里面的密码 rePwd 用户输入的密码
func ComparePassword(dbPwd, reqPwd string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(reqPwd))
	return e == nil
}
