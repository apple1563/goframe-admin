package xotp

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateOtp(uname string, issuer string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer, // 发行者名称（用于显示在谷歌验证器中）
		AccountName: uname,
	})
	return key, err
}

func ValidateOtp(totpCode string, keySecret string) bool {
	isValid := totp.Validate(totpCode, keySecret)
	return isValid
}
