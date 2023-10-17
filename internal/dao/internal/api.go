// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApiDao is the data access object for table api.
type ApiDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns ApiColumns // columns contains all the column names of Table for convenient usage.
}

// ApiColumns defines and stores column names for table api.
type ApiColumns struct {
	Id        string //
	Url       string //
	Method    string //
	Group     string //
	Remark    string //
	CreatedAt string //
	UpdatedAt string //
}

// apiColumns holds the columns for table api.
var apiColumns = ApiColumns{
	Id:        "id",
	Url:       "url",
	Method:    "method",
	Group:     "group",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewApiDao creates and returns a new DAO object for table data access.
func NewApiDao() *ApiDao {
	return &ApiDao{
		group:   "default",
		table:   "api",
		columns: apiColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApiDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApiDao) Columns() ApiColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApiDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApiDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
