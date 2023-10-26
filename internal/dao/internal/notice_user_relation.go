// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeUserRelationDao is the data access object for table notice_user_relation.
type NoticeUserRelationDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns NoticeUserRelationColumns // columns contains all the column names of Table for convenient usage.
}

// NoticeUserRelationColumns defines and stores column names for table notice_user_relation.
type NoticeUserRelationColumns struct {
	Id        string //
	NoticeId  string //
	Uid       string //
	Status    string // 1未读2已读3隐藏，用户看过把status置为2，看完后选择删除就把status置为3
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// noticeUserRelationColumns holds the columns for table notice_user_relation.
var noticeUserRelationColumns = NoticeUserRelationColumns{
	Id:        "id",
	NoticeId:  "notice_id",
	Uid:       "uid",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewNoticeUserRelationDao creates and returns a new DAO object for table data access.
func NewNoticeUserRelationDao() *NoticeUserRelationDao {
	return &NoticeUserRelationDao{
		group:   "default",
		table:   "notice_user_relation",
		columns: noticeUserRelationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoticeUserRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoticeUserRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoticeUserRelationDao) Columns() NoticeUserRelationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoticeUserRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoticeUserRelationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoticeUserRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
