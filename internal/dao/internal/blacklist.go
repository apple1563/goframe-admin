// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlacklistDao is the data access object for table blacklist.
type BlacklistDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns BlacklistColumns // columns contains all the column names of Table for convenient usage.
}

// BlacklistColumns defines and stores column names for table blacklist.
type BlacklistColumns struct {
	Id        string // 黑名单ID
	Ip        string // IP地址
	Remark    string // 备注
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// blacklistColumns holds the columns for table blacklist.
var blacklistColumns = BlacklistColumns{
	Id:        "id",
	Ip:        "ip",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewBlacklistDao creates and returns a new DAO object for table data access.
func NewBlacklistDao() *BlacklistDao {
	return &BlacklistDao{
		group:   "default",
		table:   "blacklist",
		columns: blacklistColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlacklistDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlacklistDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlacklistDao) Columns() BlacklistColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlacklistDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlacklistDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlacklistDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
