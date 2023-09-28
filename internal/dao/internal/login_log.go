// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LoginLogDao is the data access object for table login_log.
type LoginLogDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns LoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// LoginLogColumns defines and stores column names for table login_log.
type LoginLogColumns struct {
	Id          string //
	Uid         string //
	Username    string //
	Ip          string //
	CreatedAt   string //
	UpdatedAt   string //
	ClientAgent string // 注册clientAgen头
	Role        string // 1用户2代理3管理
	PRole       string //
	Pid         string //
	PUsername   string //
}

// loginLogColumns holds the columns for table login_log.
var loginLogColumns = LoginLogColumns{
	Id:          "id",
	Uid:         "uid",
	Username:    "username",
	Ip:          "ip",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	ClientAgent: "client_agent",
	Role:        "role",
	PRole:       "p_role",
	Pid:         "pid",
	PUsername:   "p_username",
}

// NewLoginLogDao creates and returns a new DAO object for table data access.
func NewLoginLogDao() *LoginLogDao {
	return &LoginLogDao{
		group:   "default",
		table:   "login_log",
		columns: loginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LoginLogDao) Columns() LoginLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
