// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ButtonDao is the data access object for table button.
type ButtonDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ButtonColumns // columns contains all the column names of Table for convenient usage.
}

// ButtonColumns defines and stores column names for table button.
type ButtonColumns struct {
	Id        string //
	MenuId    string // 按钮所在菜单id
	MenuTitle string // 按钮所在菜单名称
	Name      string // 按钮标识符
	Title     string // 按钮名称
	Remark    string //
}

// buttonColumns holds the columns for table button.
var buttonColumns = ButtonColumns{
	Id:        "id",
	MenuId:    "menu_id",
	MenuTitle: "menu_title",
	Name:      "name",
	Title:     "title",
	Remark:    "remark",
}

// NewButtonDao creates and returns a new DAO object for table data access.
func NewButtonDao() *ButtonDao {
	return &ButtonDao{
		group:   "default",
		table:   "button",
		columns: buttonColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ButtonDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ButtonDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ButtonDao) Columns() ButtonColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ButtonDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ButtonDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ButtonDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
