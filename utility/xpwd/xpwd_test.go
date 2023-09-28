package xpwd

import "testing"

func TestGenPwd(t *testing.T) {
	pwd := GenPwd("2")
	println(pwd)
}
