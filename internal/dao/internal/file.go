// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FileDao is the data access object for table file.
type FileDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns FileColumns // columns contains all the column names of Table for convenient usage.
}

// FileColumns defines and stores column names for table file.
type FileColumns struct {
	Id         string //
	Drive      string // 上传驱动：本地，ucloud
	OriginName string // 文件原始名
	FileUrl    string //
	Size       string // 单位M
	Ext        string // 扩展名
	Uid        string // 上传者用户id
	CreatedAt  string //
	UpdatedAt  string //
	Remark     string //
}

// fileColumns holds the columns for table file.
var fileColumns = FileColumns{
	Id:         "id",
	Drive:      "drive",
	OriginName: "origin_name",
	FileUrl:    "file_url",
	Size:       "size",
	Ext:        "ext",
	Uid:        "uid",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Remark:     "remark",
}

// NewFileDao creates and returns a new DAO object for table data access.
func NewFileDao() *FileDao {
	return &FileDao{
		group:   "default",
		table:   "file",
		columns: fileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FileDao) Columns() FileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
