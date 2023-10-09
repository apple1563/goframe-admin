// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserRelation is the golang structure for table user_relation.
type UserRelation struct {
	Id      int  `json:"id"      description:""`
	PUserId int  `json:"pUserId" description:""`
	Level   uint `json:"level"   description:""`
	UserId  uint `json:"userId"  description:""`
}
