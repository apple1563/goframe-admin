// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id          uint        `json:"id"          description:""`
	Username    string      `json:"username"    description:""`
	Password    string      `json:"password"    description:""`
	Nickname    string      `json:"nickname"    description:""`
	Email       string      `json:"email"       description:""`
	Phone       string      `json:"phone"       description:""`
	Status      int         `json:"status"      description:"1正常2禁用3注销"`
	ClientAgent string      `json:"clientAgent" description:"注册clientAgen头"`
	Ip          string      `json:"ip"          description:"IP"`
	RoleId      int         `json:"roleId"      description:"1用户3代理2管理"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
	PRoleId     int         `json:"pRoleId"     description:""`
	Pid         int         `json:"pid"         description:""`
	PUsername   string      `json:"pUsername"   description:""`
	RoleName    string      `json:"roleName"    description:""`
	PRoleName   string      `json:"pRoleName"   description:""`
}
