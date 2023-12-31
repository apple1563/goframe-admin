// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-starter/internal/dao/internal"
)

// internalNoticeUserRelationDao is internal type for wrapping internal DAO implements.
type internalNoticeUserRelationDao = *internal.NoticeUserRelationDao

// noticeUserRelationDao is the data access object for table notice_user_relation.
// You can define custom methods on it to extend its functionality as you wish.
type noticeUserRelationDao struct {
	internalNoticeUserRelationDao
}

var (
	// NoticeUserRelation is globally public accessible object for table notice_user_relation operations.
	NoticeUserRelation = noticeUserRelationDao{
		internal.NewNoticeUserRelationDao(),
	}
)

// Fill with you ideas below.
