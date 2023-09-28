// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LoginLog is the golang structure for table login_log.
type LoginLog struct {
	Id          int         `json:"id"          description:""`
	Uid         int         `json:"uid"         description:""`
	Username    string      `json:"username"    description:""`
	Ip          string      `json:"ip"          description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
	ClientAgent string      `json:"clientAgent" description:"注册clientAgen头"`
	Role        int         `json:"role"        description:"1用户2代理3管理"`
	PRole       string      `json:"pRole"       description:""`
	Pid         int         `json:"pid"         description:""`
	PUsername   string      `json:"pUsername"   description:""`
}
