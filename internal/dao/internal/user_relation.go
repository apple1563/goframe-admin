// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserRelationDao is the data access object for table user_relation.
type UserRelationDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns UserRelationColumns // columns contains all the column names of Table for convenient usage.
}

// UserRelationColumns defines and stores column names for table user_relation.
type UserRelationColumns struct {
	Id      string //
	PUserId string //
	Level   string //
	UserId  string //
}

// userRelationColumns holds the columns for table user_relation.
var userRelationColumns = UserRelationColumns{
	Id:      "id",
	PUserId: "p_user_id",
	Level:   "level",
	UserId:  "user_id",
}

// NewUserRelationDao creates and returns a new DAO object for table data access.
func NewUserRelationDao() *UserRelationDao {
	return &UserRelationDao{
		group:   "default",
		table:   "user_relation",
		columns: userRelationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserRelationDao) Columns() UserRelationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserRelationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
