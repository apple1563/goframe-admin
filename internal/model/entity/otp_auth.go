// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// OtpAuth is the golang structure for table otp_auth.
type OtpAuth struct {
	Id       int    `json:"id"       description:""`
	Uid      uint   `json:"uid"      description:""`
	Username string `json:"username" description:""`
	Secret   string `json:"secret"   description:""`
	Status   uint   `json:"status"   description:"1开2关"`
}
